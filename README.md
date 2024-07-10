# 🤖 AI-Powered Infrastructure Optimizer

An intelligent tool to optimize your multi-cloud infrastructure using machine learning predictions.

## 🌟 Features

- 📊 Collects resource usage data from multiple cloud providers (AWS, GCP, Azure)
- 💾 Stores historical data in a local SQLite database
- 🧠 Uses machine learning to predict future resource usage
- 💡 Generates optimization recommendations based on current and predicted usage
- 🔄 Runs periodically to provide up-to-date insights

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or later
- SQLite
- AWS, GCP, and Azure accounts with appropriate permissions

### 🛠️ Installation

1. Clone the repository:

```

git clone https://github.com/kamran0812/ai-infra-optimizer.git
cd ai-infra-optimizer

```

2. Install dependencies:

```

go mod tidy

```

3. Set up your cloud provider credentials as environment variables:

```

export AWS_ACCESS_KEY_ID=your_aws_access_key
export AWS_SECRET_ACCESS_KEY=your_aws_secret_key
export GOOGLE_APPLICATION_CREDENTIALS=path/to/your/gcp-credentials.json
export AZURE_SUBSCRIPTION_ID=your_azure_subscription_id

```

### 🏃‍♂️ Running the Optimizer

1. Build the project:

```

go build -o optimizer ./cmd/optimizer

```

2. Run the optimizer:

```

./optimizer AWS


```

OR Multiple Providers

```

./optimizer AWS AZURE GCP


```

## 🛠️ Customization

- Adjust the prediction time horizon in `optimizer.go`
- Modify the recommendation thresholds in `optimizer.go`
- Add support for additional cloud providers in the `cloud/` directory

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 🔮 Upcoming Features

We're constantly working to improve the AI-Powered Infrastructure Optimizer. Here are some exciting features on our roadmap:

- 🔷 Azure Support: Expand our cloud coverage to include Microsoft Azure, allowing for comprehensive multi-cloud optimization.

- 🔶 GCP Support: Integrate with Google Cloud Platform to provide even more flexibility in cloud resource management.

- 🧠 Updated Custom ML Model: Enhance our prediction capabilities with a more sophisticated machine learning model, improving the accuracy of our optimization recommendations.

- 📊 Advanced Visualization Dashboard: Implement a user-friendly web interface to display historical data, predictions, and recommendations in an intuitive, graphical format.

- 🔄 Real-time Optimization: Move beyond periodic checks to continuous, real-time monitoring and optimization of your cloud resources.

Stay tuned for these updates! We're excited to bring you even more powerful infrastructure optimization capabilities.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [AWS SDK for Go](https://github.com/aws/aws-sdk-go)
- [Sajari Regression Library](https://github.com/sajari/regression)
- [SQLite Driver for Go](https://github.com/mattn/go-sqlite3)
