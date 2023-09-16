import json

# Load the Nikto JSON data
with open('nikto_.json') as f:
    nikto_data = json.load(f)

# Create a basic SARIF object
sarif_data = {
    "$schema": "https://schemastore.azurewebsites.net/schemas/json/sarif-2.1.0-rtm.5.json",
    "version": "2.1.0",
    "runs": [
        {
            "tool": {
                "driver": {
                    "name": "Nikto",
                    "version": "2.1.6",
                    "informationUri": "https://cirt.net/Nikto2",
                    "rules": []
                }
            },
            "results": []
        }
    ]
}

# Convert Nikto data to SARIF format
# This is a simplified example and you would need to modify it
# according to your actual Nikto JSON output and the SARIF output you want
for item in nikto_data:
    # Add a rule for each item
    rule = {
        "id": item['id'],
        "name": item['name'],
        "shortDescription": {
            "text": item['description']
        },
        "helpUri": item['helpUri'],
        "help": {
            "text": item['help']
        }
    }
    sarif_data['runs'][0]['tool']['driver']['rules'].append(rule)

    # Add a result for each item
    result = {
        "ruleId": item['id'],
        "message": {
            "text": item['message']
        },
        "locations": [
            {
                "physicalLocation": {
                    "artifactLocation": {
                        "uri": item['uri']
                    }
                }
            }
        ]
    }
    sarif_data['runs'][0]['results'].append(result)

# Save the SARIF data to a file
with open('nikto_output.sarif', 'w') as f:
    json.dump(sarif_data, f, indent=4)
