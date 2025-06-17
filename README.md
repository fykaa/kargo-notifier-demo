# Kargo Notify Step Demo

This repository demonstrates a minimal Kargo workflow with a focus on the `notify` step, which sends notifications to a Slack channel during the promotion process.

## Prerequisites

- Kubernetes cluster
- Kargo controller installed
- Slack API token and channel ID

## Usage

1. Update secrets with your Git and Slack credentials.
2. Apply all manifests in the `slack-notify-step-demo` directory:
   ```sh
   kubectl apply -f slack-notify-step-demo/
   ```
3. Trigger a promotion to see the notification in Slack.
