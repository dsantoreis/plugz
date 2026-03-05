from __future__ import annotations

from dataclasses import dataclass
from .rag import SimpleRAG


@dataclass
class WorkflowResult:
    plan: str
    context: str
    draft: str
    review: str


class PlannerAgent:
    def run(self, goal: str) -> str:
        return (
            f"Goal: {goal}\n"
            "Plan:\n"
            "1) Retrieve relevant evidence\n"
            "2) Draft practical recommendations\n"
            "3) Review for clarity and risk"
        )


class ResearcherAgent:
    def __init__(self, rag: SimpleRAG):
        self.rag = rag

    def run(self, goal: str) -> str:
        return self.rag.context_block(goal, top_k=3)


class WriterAgent:
    def run(self, goal: str, context: str) -> str:
        bullets = [
            "- Focus on first-time user time-to-value.",
            "- Keep steps short and explicit.",
            "- Add in-product help near complex actions.",
            "- Track activation metrics weekly.",
        ]
        return (
            f"Question: {goal}\n\n"
            "Recommended approach:\n"
            + "\n".join(bullets)
            + f"\n\nGrounding context:\n{context}"
        )


class ReviewerAgent:
    def run(self, draft: str) -> str:
        notes = []
        if "metrics" not in draft.lower():
            notes.append("Missing success metrics.")
        if len(draft) < 200:
            notes.append("Draft may be too short for stakeholders.")
        if not notes:
            notes.append("Looks clear and actionable.")
        return "Review notes: " + " ".join(notes)


class MultiAgentOrchestrator:
    def __init__(self, rag: SimpleRAG):
        self.planner = PlannerAgent()
        self.researcher = ResearcherAgent(rag)
        self.writer = WriterAgent()
        self.reviewer = ReviewerAgent()

    def run(self, goal: str) -> WorkflowResult:
        plan = self.planner.run(goal)
        context = self.researcher.run(goal)
        draft = self.writer.run(goal, context)
        review = self.reviewer.run(draft)
        return WorkflowResult(plan=plan, context=context, draft=draft, review=review)
