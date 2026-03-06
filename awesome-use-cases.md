# Awesome Plugz Use Cases

Ten practical ways teams can use Plugz in production.

## 1. Support triage routing
Problem: inbound tickets land in one queue and first response takes too long.
Solution: register a `triage-router` skill that classifies severity and team ownership.
Result: faster first assignment and cleaner SLA tracking.

```bash
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Authorization: Bearer dev-token" \
  -H "Content-Type: application/json" \
  -d '{"skill":"triage-router","input":{"subject":"Checkout failed","body":"Customer reports card declined"}}'
```

## 2. Lead enrichment pipeline
Problem: SDRs manually enrich lead records from fragmented sources.
Solution: compose `company-enrich` and `persona-score` skills.
Result: richer lead context in minutes instead of hours.

## 3. Incident command assistant
Problem: on-call engineers lose time collecting context during incidents.
Solution: run `incident-brief` skill to summarize alerts, recent deploys, and known runbooks.
Result: shorter time-to-diagnosis for Sev-2 and Sev-1 events.

## 4. RAG content refresh
Problem: knowledge bases drift and retrieval quality degrades.
Solution: schedule `kb-refresh` skill to re-chunk, re-embed, and validate freshness.
Result: more accurate answers and fewer stale references.

## 5. Compliance evidence packaging
Problem: audits require repeated manual export of access and change logs.
Solution: execute `audit-bundle` skill that gathers artifacts from approved systems.
Result: repeatable evidence packs per control and period.

## 6. Release notes generation
Problem: release notes are delayed and inconsistent.
Solution: use `release-digest` skill over merged PRs and linked issues.
Result: predictable release communication every sprint.

## 7. Customer onboarding checklist automation
Problem: onboarding tasks are scattered across teams.
Solution: orchestrate `tenant-bootstrap` skill to create project defaults and required integrations.
Result: lower onboarding variance and faster time-to-value.

## 8. Localization workflow
Problem: product copy updates stall because translation handoff is manual.
Solution: run `content-localize` skill with locale-specific style rules.
Result: faster multi-language releases with fewer review rounds.

## 9. Data quality guardrails
Problem: downstream dashboards break on malformed upstream payloads.
Solution: apply `schema-guard` skill before persistence.
Result: fewer broken reports and cleaner analytics datasets.

## 10. Sales proposal drafting
Problem: proposal creation is slow and repetitive for similar deal shapes.
Solution: trigger `proposal-draft` skill with account profile and scope.
Result: quicker first draft and consistent structure across deals.

---

If you built a useful pattern with Plugz, open an issue and share your workflow so others can reproduce it.
