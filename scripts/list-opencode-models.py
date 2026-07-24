#!/usr/bin/env python3
"""List all available models from OpenCode Zen

Requirements:
    pip install requests python-dotenv

Environment:
    Set ANTHROPIC_API_KEY (your OpenCode Zen API key)
"""

from dotenv import load_dotenv
import os
import requests
import sys
import json

load_dotenv()

api_key = os.getenv('ANTHROPIC_API_KEY')
if not api_key:
    print("Error: ANTHROPIC_API_KEY environment variable not set")
    sys.exit(1)

endpoint = "https://opencode.ai/zen/v1/models"

try:
    response = requests.get(
        endpoint,
        headers={'x-api-key': api_key}
    )

    if response.status_code == 200:
        data = response.json()
        models = data.get('data', [])

        print("\n" + "=" * 70)
        print("OpenCode Zen - Available Models")
        print("=" * 70 + "\n")

        if models:
            for model in models:
                model_id = model.get('id', 'unknown')
                model_name = model.get('name', model_id)
                owner = model.get('owned_by', 'unknown')

                print(f"ID: {model_id}")
                print(f"  Name: {model_name}")
                print(f"  Owner: {owner}")
                print()
        else:
            print("No models found")

        print("=" * 70)
        print(f"Total models: {len(models)}\n")

    else:
        print(f"Error: HTTP {response.status_code}")
        print(f"Response: {response.text}")
        sys.exit(1)

except requests.exceptions.RequestException as e:
    print(f"Error: Failed to connect to OpenCode Zen")
    print(f"Details: {e}")
    sys.exit(1)
