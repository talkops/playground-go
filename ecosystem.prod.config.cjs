module.exports = {
  apps: [
    {
      cmd: 'talkops',
      name: 'client',
    },
    {
      autorestart: true,
      cmd: 'go build -o /tmp/app src/main.go && /tmp/app',
      error_file: process.env.TALKOPS_STDERR,
      name: 'extension',
      out_file: process.env.TALKOPS_STDOUT,
    },
  ],
}
