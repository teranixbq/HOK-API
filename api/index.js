const { exec } = require('child_process');

// Jalankan file Go
const server = exec('build/app', (error, stdout, stderr) => {
    if (error) {
        console.error(`exec error: ${error}`);
        return;
    }
    console.log(`stdout: ${stdout}`);
    console.error(`stderr: ${stderr}`);
});

// Tangani penghentian proses
process.on('SIGINT', () => {
    server.kill('SIGINT');
    process.exit();
});
