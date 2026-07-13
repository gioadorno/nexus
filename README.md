# AutoGen FDE Workflow Engine

> **Portfolio Highlight:** This repository demonstrates Forward Deployed Engineering (FDE) principles utilizing **Microsoft AutoGen**. It simulates a multi-agent AI workflow that accelerates the discovery-to-delivery lifecycle by automating the conversion of raw meeting notes into structured product specs, architectural reviews, and ready-to-work engineering tickets.

## Features

1. **Dynamic Input/Output:** Point the script at any Markdown file containing meeting notes. It will automatically generate a clean `[File_Name]_FDE_Plan.md` file in the same directory.
2. **Human-in-the-Loop:** The script pauses before generating tickets, allowing the human FDE to inject feedback or correct the architecture.
3. **Linear API Tool with Fallback:** The Ticket Agent calls a Python function to create tickets in Linear. If the API key is missing or fails, it gracefully falls back to appending the tickets directly into your local Markdown plan.
4. **Hybrid Semantic Memory (ChromaDB):** Agents learn from past executions using a local ChromaDB vector database. A post-execution hook extracts these learned rules and writes them to a visible `fde_knowledge_base.md` file so you can audit the AI's "brain."

## How to Use It

### 1. Setup Environment
Because this relies on specific library versions, run it from the provided virtual environment:
```bash
cd ~/fde-autogen
source venv/bin/activate
```

### 2. Configure Your LLM
Export your OpenAI API Key (or edit `fde_workflow.py` to point to a local Ollama model / Vertex Gemini):
```bash
export OPENAI_API_KEY="sk-your-actual-api-key"
```

### 3. Run the Engine
Pass the path to your meeting notes as an argument:
```bash
python fde_workflow.py ~/Obsidian/Meetings/Acme_Kickoff.md
```

Watch the terminal as the Discovery Agent, Architect Agent, and Ticket Agent debate and build your plan!
