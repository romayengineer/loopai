# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Purpose

This repository is a learning project focused on understanding and documenting AI agents, agentic loops, and workflow automation patterns. The goal is to:
1. Document theoretical concepts and patterns for building AI agents
2. Explore practical implementations of agent frameworks and workflows
3. Experiment with multi-agent orchestration and autonomous systems
4. Track learnings and implementations as they evolve

## Repository Structure

The repository is organized into learning phases:

- **`docs/`** – Conceptual documentation on AI agents, loops, and workflows. Start here for theoretical understanding.
- **`examples/`** – Working code examples demonstrating specific patterns (single agents, loops, multi-agent workflows).
- **`experiments/`** – Sandbox for trying new approaches and testing ideas before formalizing them.
- **`references/`** – Links, summaries, and notes from external resources.

## Development Workflow

Since this is a learning repository:
1. **Documentation First** – Write concepts and patterns in `docs/` before implementing
2. **Example-Driven** – Each pattern should have a working example in `examples/`
3. **Experiment Freely** – Use `experiments/` to test ideas without affecting main documentation
4. **Iterate** – Move proven experiments into `examples/` and refine documentation as understanding deepens

## Key Concepts to Document

As you explore, focus on:
- **Agent Fundamentals** – What makes an AI agent, decision-making loops, state management
- **Agentic Loops** – Tool calling patterns, reasoning loops, feedback mechanisms
- **Workflow Patterns** – Sequential agents, parallel agents, conditional routing, state passing
- **Multi-Agent Orchestration** – Agent coordination, communication patterns, supervision
- **Autonomy & Safety** – Agent constraints, guardrails, error handling in autonomous systems

## Guidelines for Future Work

- Keep documentation and code close – examples should directly illustrate their corresponding concepts
- Use clear naming conventions: use `agent-` prefix for agent-focused code, `loop-` for loop patterns, `workflow-` for orchestration examples
- When adding implementations, include brief comments explaining which conceptual pattern it demonstrates
- Track assumptions and constraints discovered during exploration – these often become important design decisions

## No Build/Test Infrastructure Yet

This is a documentation-first project. As implementations grow, standard tooling will be added (linting, testing, CI/CD). For now, focus on clarity and correctness in documentation and runnable examples.
