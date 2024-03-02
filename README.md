# SMTP-sendEmail

SMTP-sendEmail is a Golang web server that facilitates sending emails via SMTP. It includes an HTTP endpoint for sending emails and a simple HTML form to interact with the server.

## Getting Started

Follow these instructions to set up and run the SMTP-sendEmail on your local machine.

### Prerequisites

Make sure you have the following installed:

- Go (Golang): [Installation Guide](https://golang.org/doc/install)
- Git: [Installation Guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yuriy0803/SMTP-sendEmail.git
    cd SMTP-sendEmail
    ```

2. Update the `main.go` file with your SMTP server details:

    ```go
    // Update these SMTP settings with your own
    smtpHost := "your_smtp_server"
    smtpPort := 587
    smtpUsername := "your_username"
    smtpPassword := "your_password"
    ```

### Usage

1. Run the Golang server:

    ```bash
    go run main.go
    ```

2. Open your web browser and navigate to [http://localhost:3000](http://localhost:3000) to access the HTML form.

3. Fill in the form with the recipient's email address, subject, and message. Click the "Send Email" button.

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
