package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"time"
)

type Stats struct {
	Timestamp   string  `json:"timestamp"`
	Goroutines  int     `json:"goroutines"`
	MemoryAlloc float64 `json:"memory_alloc"`
	MemoryTotal float64 `json:"memory_total"`
	CPUCores    int     `json:"cpu_cores"`
}

func getStats() Stats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return Stats{
		Timestamp:   time.Now().Format("15:04:05"),
		Goroutines:  runtime.NumGoroutine(),
		MemoryAlloc: float64(m.Alloc) / 1024 / 1024,
		MemoryTotal: float64(m.TotalAlloc) / 1024 / 1024,
		CPUCores:    runtime.NumCPU(),
	}
}

func main() {
	// HTML страница
	tmpl := template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Go System Monitor</title>
    <style>
        body {
            font-family: 'Monaco', monospace;
            background: #0d1117;
            color: #58a6ff;
            padding: 20px;
            margin: 0;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
        }
        h1 {
            color: #58a6ff;
            text-align: center;
            text-shadow: 0 0 10px #58a6ff;
        }
        .stat-box {
            background: #161b22;
            border: 2px solid #30363d;
            border-radius: 8px;
            padding: 20px;
            margin: 15px 0;
            box-shadow: 0 0 20px rgba(88, 166, 255, 0.1);
        }
        .stat-label {
            color: #8b949e;
            font-size: 14px;
            margin-bottom: 5px;
        }
        .stat-value {
            color: #58a6ff;
            font-size: 32px;
            font-weight: bold;
        }
        .pulse {
            animation: pulse 2s infinite;
        }
        @keyframes pulse {
            0%, 100% { opacity: 1; }
            50% { opacity: 0.5; }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>⚡ Go Runtime Monitor ⚡</h1>
        <div class="stat-box">
            <div class="stat-label">⏰ Current Time</div>
            <div class="stat-value pulse" id="timestamp">--:--:--</div>
        </div>
        <div class="stat-box">
            <div class="stat-label">🔄 Active Goroutines</div>
            <div class="stat-value" id="goroutines">0</div>
        </div>
        <div class="stat-box">
            <div class="stat-label">💾 Memory Allocated (MB)</div>
            <div class="stat-value" id="memory">0.00</div>
        </div>
        <div class="stat-box">
            <div class="stat-label">📊 Total Memory Used (MB)</div>
            <div class="stat-value" id="total">0.00</div>
        </div>
        <div class="stat-box">
            <div class="stat-label">⚙️ CPU Cores</div>
            <div class="stat-value" id="cores">0</div>
        </div>
    </div>
    
    <script>
        function updateStats() {
            fetch('/stats')
                .then(response => response.json())
                .then(data => {
                    document.getElementById('timestamp').textContent = data.timestamp;
                    document.getElementById('goroutines').textContent = data.goroutines;
                    document.getElementById('memory').textContent = data.memory_alloc.toFixed(2);
                    document.getElementById('total').textContent = data.memory_total.toFixed(2);
                    document.getElementById('cores').textContent = data.cpu_cores;
                });
        }
        
        // Обновляем каждую секунду
        setInterval(updateStats, 1000);
        updateStats();
    </script>
</body>
</html>
    `))

	// Handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getStats())
	})

	// Создаем фоновые горутины для демонстрации
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
			}
		}()
	}

	fmt.Println("🚀 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
