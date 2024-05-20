
document.addEventListener('DOMContentLoaded', function() {
    const ctxCPU = document.getElementById('cpuChart').getContext('2d');
    const cpuChart = new Chart(ctxCPU, {
        type: 'line',
        data: {
            labels: Array.from({length: cpuData.length}, (v, k) => k + 1),
            datasets: [{
                label: 'CPU Usage (%)',
                data: cpuData,
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderColor: 'rgba(75, 192, 192, 1)',
                borderWidth: 1
            }]
        },
        options: {
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });

    const ctxMemory = document.getElementById('memoryChart').getContext('2d');
    const memoryChart = new Chart(ctxMemory, {
        type: 'doughnut',
        data: {
            labels: ['Used', 'Free'],
            datasets: [{
                label: 'Memory Usage (%)',
                data: [memoryData, 100 - memoryData],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)'
                ],
                borderWidth: 1
            }]
        },
        options: {
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
});
