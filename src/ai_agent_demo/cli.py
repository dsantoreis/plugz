from __future__ import annotations

import argparse
from pathlib import Path

from .agents import MultiAgentOrchestrator
from .rag import SimpleRAG


def build_parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(description="Run multi-agent + RAG demo")
    p.add_argument("goal", help="Goal/question for the agents")
    p.add_argument(
        "--kb",
        default=str(Path(__file__).resolve().parents[2] / "data" / "knowledge_base.txt"),
        help="Path to knowledge base text file",
    )
    return p


def main() -> None:
    args = build_parser().parse_args()
    rag = SimpleRAG.from_file(args.kb)
    orchestrator = MultiAgentOrchestrator(rag)
    result = orchestrator.run(args.goal)

    print("=== PLAN ===")
    print(result.plan)
    print("\n=== CONTEXT ===")
    print(result.context)
    print("\n=== DRAFT ===")
    print(result.draft)
    print("\n=== REVIEW ===")
    print(result.review)


if __name__ == "__main__":
    main()
