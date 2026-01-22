package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const TOTAL_TIME_FOR_SYSTEM = 120

type Booking struct {
	PersonID  int
	RoomID    int
	StartTime int
	EndTime   int
}
type MeetingRoom struct {
	ID       int
	Name     string
	mu       sync.Mutex
	bookings map[int]*Booking //key is booking start time
}
type MeetingRoomManager struct {
	rooms map[int]*MeetingRoom //key is room ID
	mu    sync.RWMutex
}
type Request struct {
	Type      string //B or C
	PersonID  int
	RoomID    int
	RequestAt int
	StartTime int
	EndTime   int
}

func NewMeetingRoomManager() *MeetingRoomManager {
	manager := &MeetingRoomManager{
		rooms: make(map[int]*MeetingRoom),
	}
	for i := 1; i <= 5; i++ {
		manager.rooms[i] = &MeetingRoom{ID: i, Name: fmt.Sprintf("Room-%d", i), bookings: make(map[int]*Booking)}
	}
	return manager
}

func ProcessRequests(manager *MeetingRoomManager, requests []Request) {
	startTime := time.Now()
	var wg sync.WaitGroup

	requestChan := make(chan Request, len(requests))

	// worker pool abt to strt
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for req := range requestChan {
				if req.Type == "B" {
					booking := &Booking{
						PersonID:  req.PersonID,
						RoomID:    req.RoomID,
						StartTime: req.StartTime,
						EndTime:   req.EndTime,
					}
					manager.Book(booking)
				} else if req.Type == "C" {
					manager.Cancel(req.PersonID, req.RoomID)
				}
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, req := range requests {
			elapsed := int(time.Since(startTime).Seconds())
			if req.RequestAt > elapsed {
				time.Sleep(time.Duration(req.RequestAt-elapsed) * time.Second)
			}
			requestChan <- req
		}
		close(requestChan)
	}()
	//
	wg.Wait()
}
func (m *MeetingRoomManager) Book(b *Booking) bool {
	m.mu.RLock() // read lock
	room, exists := m.rooms[b.RoomID]
	m.mu.RUnlock()

	if !exists {
		fmt.Printf("Room %d does not exist\n", b.RoomID)
		return false
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	for _, existing := range room.bookings {
		if (b.StartTime < existing.EndTime) && (existing.StartTime < b.EndTime) {
			fmt.Printf("Booking conflict for Room %d: %v conflicts with existing booking %v\n", b.RoomID, b, existing)
			return false
		}
	}

	room.bookings[b.StartTime] = b
	fmt.Printf("Booking successful: %v\n", b)
	return true
}

func (m *MeetingRoomManager) Cancel(personID, roomID int) {
	m.mu.RLock()
	room, exists := m.rooms[roomID]
	m.mu.RUnlock()

	if !exists {
		fmt.Printf("Room %d does not exist\n", roomID)
		return
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	for key, booking := range room.bookings {
		if booking.PersonID == personID {
			delete(room.bookings, key)
			fmt.Printf("Cancellation successful: %v\n", booking)
			return
		}
	}

	fmt.Printf("No booking found for Person %d in Room %d to cancel\n", personID, roomID)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var requests []Request
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, ",")

		typ := parts[0]
		personID, _ := strconv.Atoi(parts[1])
		roomID, _ := strconv.Atoi(parts[2])
		requestAt, _ := strconv.Atoi(parts[3])
		startTime, _ := strconv.Atoi(parts[4])
		endTime, _ := strconv.Atoi(parts[5])
		requests = append(requests, Request{typ, personID, roomID, requestAt, startTime, endTime})
	}

	fmt.Println("Meeting Room Booking System Initialized")
	fmt.Println("------------------------------------")

	sort.Slice(requests, func(i, j int) bool {
		return requests[i].RequestAt < requests[j].RequestAt
	})

	manager := NewMeetingRoomManager()
	ProcessRequests(manager, requests)

	fmt.Println("------------------------------------")
	fmt.Println("Meeting Room Booking System Terminated")

}
