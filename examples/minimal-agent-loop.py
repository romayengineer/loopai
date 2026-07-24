"""Minimal AI Agent Loop with LangChain

This is the simplest possible agent loop:
- An LLM (Claude Haiku 4.5 via OpenCode Zen)
- Simple tools (arithmetic)
- An agentic loop that repeats until the goal is achieved

OpenCode Zen provides access to 50+ models (Anthropic, OpenAI, Qwen, etc.) with a single API key.

Requirements:
    pip install langchain langchain-anthropic anthropic python-dotenv

Environment:
    Create a .env file with: ANTHROPIC_API_KEY=your-opencode-zen-key
"""

from dotenv import load_dotenv
from langchain_anthropic import ChatAnthropic
from langchain_core.tools import tool
from langchain_core.messages import HumanMessage, ToolMessage

load_dotenv()

# Initialize the LLM (using Haiku 4.5 via OpenCode Zen - 50+ models with one API key)
llm = ChatAnthropic(
    model="claude-haiku-4-5",
    temperature=0,
    anthropic_api_url="https://opencode.ai/zen"
)

# Define simple tools
@tool
def add(a: int, b: int) -> int:
    """Add two numbers."""
    return a + b

@tool
def multiply(a: int, b: int) -> int:
    """Multiply two numbers."""
    return a * b

tools = [add, multiply]

# Bind tools to the LLM
llm_with_tools = llm.bind_tools(tools)

# Run the agent loop
messages = [HumanMessage(content="What is 25 * 4 + 10?")]

print("Starting agent loop...")
print(f"User: {messages[0].content}\n")

max_iterations = 10
iteration = 0

while iteration < max_iterations:
    iteration += 1

    # Get LLM response
    response = llm_with_tools.invoke(messages)
    messages.append(response)

    # Check if we're done (no tool calls)
    if not response.tool_calls:
        print(f"\nFinal Answer: {response.content}")
        break

    # Process tool calls
    for tool_call in response.tool_calls:
        tool_name = tool_call["name"]
        tool_args = tool_call["args"]

        # Call the tool by name
        tool_map = {"add": add, "multiply": multiply}
        tool_func = tool_map[tool_name]
        result = tool_func.invoke(tool_args)

        print(f"Tool call: {tool_name}({tool_args}) = {result}")

        # Add tool result to messages
        messages.append(ToolMessage(
            content=str(result),
            tool_call_id=tool_call["id"]
        ))

if iteration >= max_iterations:
    print(f"\nReached max iterations ({max_iterations})")
