# AI Agent Skills Demo (Multi-Agent + RAG)

Private demo repository showcasing:

- **Multi-agent orchestration** (planner, researcher, writer, reviewer)
- **Lightweight RAG pipeline** (local docs + lexical retrieval)
- **Executable CLI** to run end-to-end workflows
- **Minimal tests** for orchestration and retrieval

## Quick start

```bash
cd ~/Projects/ai-agent-skills-demo
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python -m ai_agent_demo.cli "How can we improve customer onboarding?"
```

## Example output

The CLI prints:
1. plan
2. retrieved context snippets
3. final draft
4. reviewer notes

## Run tests

```bash
python -m unittest discover -s tests -p 'test_*.py' -v
```

## Project structure

- `src/ai_agent_demo/agents.py` — agent roles and orchestration logic
- `src/ai_agent_demo/rag.py` — document chunking + retrieval
- `src/ai_agent_demo/cli.py` — runnable CLI entrypoint
- `tests/` — minimal regression coverage

