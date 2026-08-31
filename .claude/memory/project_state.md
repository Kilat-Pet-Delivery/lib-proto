---
name: project_state
description: The resume point for this repo - current checkpoint (sha, environment, open units table, recommended next unit) at the top, earlier checkpoints below. Read first in every session; rewritten by /recap.
metadata:
  type: project
---

## 2026-08-31 state (resume here)

- **Repo:** `main` @ `231585e` - feat(dto): add shop + inventory DTOs; extend BookingStatus and Booking
- **Environment:** dev-infra stack up (`./dev.ps1 up kilat`).
- **Open units**

| Unit / ticket | State | Blocked on | Note |
|---|---|---|---|
| KPD-63 gofmt sweep | In Review | review | PR #4 |
| KPD-52 CHANGELOG | In Review | review | PR #5 |
| KPD-53 CI caller workflow | In Progress | **a credential scope** | committed locally, push blocked on the workflow scope |

- **Recommended next unit:** nothing specific - this repo moves when an event or DTO changes.
- **Waiting on Luqman:** merge the open PRs above. Several are stacked, so order matters.

## Earlier checkpoints

(none - this layer was created 2026-08-31 under KPD-51)
