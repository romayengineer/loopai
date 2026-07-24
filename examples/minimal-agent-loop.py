"""Minimal AI Agent Loop with LangChain

This is the simplest possible agent loop:
- An LLM (Claude or GPT-4)
- One simple tool (arithmetic)
- An agentic loop that repeats until the goal is achieved

Requirements:
    pip install langchain anthropic

Environment:
    Set ANTHROPIC_API_KEY or OPENAI_API_KEY
"""

from langchain_anthropic import ChatAnthropic
from langchain.agents import tool, AgentExecutor, create_tool_calling_agent
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder

# Initialize the LLM
llm = ChatAnthropic(model="claude-3-5-sonnet-20241022", temperature=0)

# Define a simple tool
@tool
def add(a: int, b: int) -> int:
    """Add two numbers."""
    return a + b

@tool
def multiply(a: int, b: int) -> int:
    """Multiply two numbers."""
    return a * b

tools = [add, multiply]

# Create the agent prompt
prompt = ChatPromptTemplate.from_messages([
    ("system", "You are a helpful math assistant. Use tools to solve problems."),
    ("human", "{input}"),
    MessagesPlaceholder(variable_name="agent_scratchpad"),
])

# Create and run the agent
agent = create_tool_calling_agent(llm, tools, prompt)
executor = AgentExecutor(agent=agent, tools=tools, verbose=True)

# Run the agent loop
result = executor.invoke({"input": "What is 25 * 4 + 10?"})
print("\n" + "="*50)
print("FINAL ANSWER:", result["output"])
