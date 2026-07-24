# Minimal Agent Loop with LangChain

This example demonstrates the simplest possible AI agent loop using LangChain.

## What This Does

Given a math problem like "What is 25 * 4 + 10?", the agent:
1. Reads the input
2. Decides which tool(s) to use (add, multiply)
3. Calls the tools
4. Gets results back
5. Repeats until it can answer the question
6. Returns the final answer

## Code Breakdown

```python
from langchain_anthropic import ChatAnthropic
from langchain.agents import tool, AgentExecutor, create_tool_calling_agent
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
```
Import LangChain components for agents, tools, and prompts.

```python
llm = ChatAnthropic(
    model="claude-haiku-4-5",
    temperature=0,
    anthropic_api_url="https://opencode.ai/zen/v1"
)
```
Initialize Claude Haiku 4.5 via OpenCode Zen. OpenCode Zen provides access to 50+ models (Anthropic, OpenAI, Qwen, etc.) with a single API key, making it ideal for learning and experimenting with different models.

```python
@tool
def add(a: int, b: int) -> int:
    """Add two numbers."""
    return a + b
```
Define tools with the `@tool` decorator. The docstring describes what the tool does (used by the LLM).

```python
agent = create_tool_calling_agent(llm, tools, prompt)
executor = AgentExecutor(agent=agent, tools=tools, verbose=True)
```
Create the agent and executor. `verbose=True` shows the reasoning loop.

```python
result = executor.invoke({"input": "What is 25 * 4 + 10?"})
```
Run the agent loop. It will:
- Call `multiply(25, 4)` → 100
- Call `add(100, 10)` → 110
- Return the answer

## Running It

```bash
# Install dependencies
pip install langchain langchain-anthropic anthropic python-dotenv

# Create a .env file
echo 'ANTHROPIC_API_KEY="your-zen-api-key-here"' > .env

# Run
python minimal-agent-loop.py
```

## Key Insights

- **The loop is implicit:** `AgentExecutor` handles the "observe → decide → act" cycle
- **Tool use is automatic:** The LLM decides when to call tools based on the task
- **Minimal setup:** 50 lines of code for a fully functioning agent
- **Extensible:** Add more tools by defining more `@tool` functions

## What's Missing (Compared to Production)

- Error handling for tool failures
- Iteration limits (prevent infinite loops)
- Logging and monitoring
- Structured output
- Multi-step planning
- Memory/context management

These will be explored in subsequent examples.

## Next Steps

- Modify the tools (try adding `divide`, `subtract`)
- Change the prompt to test different reasoning styles
- Try more complex problems that require multiple tool calls
- Switch to a different LLM (GPT-4, Llama, etc.)
