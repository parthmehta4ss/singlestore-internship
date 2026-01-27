# MeetingRoomBookingSystem — Rendered Manifests Diff (local experiment)

This small folder is a local experiment to demonstrate the "rendered manifests" pattern: render the base and overlay Kubernetes manifests, then diff the two rendered outputs to show exactly what changes in Kubernetes (not the patch file).

Why: reviewers and CI should see the concrete manifest changes that will be applied to the cluster, not the Kustomize patch format.

How I did this

1. Render base
```
kustomize build k8s/base > base.yaml
```

2. Render dev overlay
```
kustomize build k8s/overlays/dev > dev.yaml
```

3. Run diff
```
diff -u base.yaml dev.yaml
```

You should see something like:
```
-spec:
-  replicas: 1
+spec:
+  replicas: 2
```

This shows the actual manifest change (`replicas: 1` -> `replicas: 2`) — exactly what a reviewer or CI should inspect.

4. (Optional) Pretty git-style diff
```
git diff --no-index base.yaml dev.yaml
```

This mimics GitHub PR diff style and is easy to read in CI.

Files in this folder

- `base.yaml` — rendered base output
- `dev.yaml` — rendered dev overlay output
- `diff.yaml` — sample output of a diff (optional)

Notes

- This is intentionally small and local — good for trying the pattern on a large project before adding CI rules.
- Use the rendered-manifests diff in PR checks to make changes crystal clear to reviewers.

Happy experimenting!
