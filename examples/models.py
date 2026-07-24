"""Pre-configured LLM models for examples."""

from langchain_anthropic import ChatAnthropic


def ModelClaudeHaiku45():
    """Claude Haiku 4.5 via OpenCode Zen - ready to use without parameters."""
    return ChatAnthropic(
        model="claude-haiku-4-5",
        temperature=0,
        anthropic_api_url="https://opencode.ai/zen",
    )
