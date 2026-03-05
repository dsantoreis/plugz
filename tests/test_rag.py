from tests._bootstrap import *
import unittest
from ai_agent_demo.rag import SimpleRAG


class TestSimpleRAG(unittest.TestCase):
    def test_retrieve_finds_relevant_chunk(self):
        text = "apple banana\n\nagent orchestration retrieval\n\nmetrics onboarding"
        rag = SimpleRAG.from_text(text, chunk_size=30)
        found = rag.retrieve("how to improve onboarding metrics", top_k=2)
        self.assertTrue(found)
        joined = " ".join(c.text for c in found).lower()
        self.assertIn("onboarding", joined)


if __name__ == "__main__":
    unittest.main()
