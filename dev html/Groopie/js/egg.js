document.addEventListener('DOMContentLoaded', function() {
    const toast = document.getElementById('toast');
    toast.classList.add('show');
});

const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

const leaves = []; // Array to store leaves

function createLeaves() {
    for (let i = 0; i < 100; i++) { // Number of leaves
        const x = Math.random() * canvas.width;
        const y = Math.random() * canvas.height;
        const size = Math.random() * 20 + 1; // Random size
        const speed = size / 10; // Random speed
        leaves.push({ x, y, size, speed });
    }
}

function drawLeaves() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    for (let i = 0; i < leaves.length; i++) {
        const leaf = leaves[i];
        ctx.fillStyle = '#9FFCF2'; // Yellow color for leaves
        ctx.beginPath();
        ctx.arc(leaf.x, leaf.y, leaf.size, 0, Math.PI * 2);
        ctx.fill();

        leaf.y += leaf.speed; // Move leaf downwards
        leaf.x += Math.sin( leaf.y / 20); // Move leaf downwards

        // Reset leaf position if it goes out of the screen
        if (leaf.y - leaf.size > canvas.height) {
            leaf.y = -leaf.size;
        }
    }
}
createLeaves();

function animate() {
    requestAnimationFrame(animate);
    drawLeaves();
}

animate();

// Redraw leaves when the window is resized
window.addEventListener('resize', function() {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    createLeaves();
});

