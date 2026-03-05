from tests._bootstrap import *
import unittest
from ai_agent_demo.agents import MultiAgentOrchestrator
from ai_agent_demo.rag import SimpleRAG


class TestOrchestrator(unittest.TestCase):
    def test_workflow_runs_end_to_end(self):
        rag = SimpleRAG.from_text("onboarding metrics activation")
        orch = MultiAgentOrchestrator(rag)
        out = orch.run("improve onboarding")
        self.assertIn("Plan:", out.plan)
        self.assertTrue(out.context)
        self.assertIn("Recommended approach", out.draft)
        self.assertIn("Review notes", out.review)


if __name__ == "__main__":
    unittest.main()
