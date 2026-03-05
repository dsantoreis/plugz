from __future__ import annotations

from dataclasses import dataclass
from pathlib import Path
import re
from typing import Iterable, List

TOKEN_RE = re.compile(r"[a-zA-Z0-9_]+")


def tokenize(text: str) -> set[str]:
    return {t.lower() for t in TOKEN_RE.findall(text)}


@dataclass
class Chunk:
    id: int
    text: str


class SimpleRAG:
    def __init__(self, chunks: List[Chunk]):
        self.chunks = chunks

    @classmethod
    def from_text(cls, text: str, chunk_size: int = 280) -> "SimpleRAG":
        lines = [ln.strip() for ln in text.splitlines() if ln.strip()]
        chunks: List[Chunk] = []
        buff = ""
        cid = 0
        for ln in lines:
            if len(buff) + len(ln) + 1 > chunk_size and buff:
                chunks.append(Chunk(cid, buff.strip()))
                cid += 1
                buff = ""
            buff += ln + "\n"
        if buff.strip():
            chunks.append(Chunk(cid, buff.strip()))
        return cls(chunks)

    @classmethod
    def from_file(cls, path: str | Path) -> "SimpleRAG":
        return cls.from_text(Path(path).read_text(encoding="utf-8"))

    def retrieve(self, query: str, top_k: int = 3) -> List[Chunk]:
        q = tokenize(query)
        scored: list[tuple[int, int, Chunk]] = []
        for c in self.chunks:
            ct = tokenize(c.text)
            overlap = len(q.intersection(ct))
            scored.append((overlap, len(c.text), c))
        scored.sort(key=lambda x: (x[0], x[1]), reverse=True)
        return [c for s, _, c in scored if s > 0][:top_k]

    def context_block(self, query: str, top_k: int = 3) -> str:
        found = self.retrieve(query, top_k=top_k)
        if not found:
            return "No relevant context found."
        return "\n\n".join(f"[chunk:{c.id}] {c.text}" for c in found)
