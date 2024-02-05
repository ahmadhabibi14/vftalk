import chokidar from 'chokidar';
import {exec} from 'child_process';
import {join} from 'path';

// Command to execute when changes occur
const commandToExecute = 'astro build';

const ignorePath = new Set( [
  'node_modules',
  '.vscode',
  '.idea',
  '.git',
  '.gitignore',
  'build.js',
  'package-lock.json',
  'package.json',
  'pnpm-lock.yaml',
  'README.md',
  'build.js',
  'dist',
] );

executeCommand(commandToExecute);

// Initialize chokidar watcher
const watcher = chokidar
  .watch( '.', {ignored: s => ignorePath.has( s ) || ignorePath.has( join( './', s ) ), ignoreInitial: true} )
  .on( 'change', (path) => executeCommand(commandToExecute) )
  .on( 'add', (path) => executeCommand(commandToExecute) )
  .on( 'unlink', (path) => executeCommand(commandToExecute) )
  .on( 'ready', () => {
    console.log( `watching files/dirs for changes` );
  } )
  .on( 'error', err => console.log( 'ERROR:', err ) );

// Function to execute command
function executeCommand(command) {
  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.error(`stderr: ${stderr}`);
  });
}

console.log('Watching for file changes...');
