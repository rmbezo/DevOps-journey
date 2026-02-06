package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>3D Octagon Sphere</title>
    <style>
        body {
            margin: 0;
            padding: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            overflow: hidden;
            font-family: Arial, sans-serif;
        }
        
        canvas {
            border: 3px solid rgba(255, 255, 255, 0.3);
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
            background: radial-gradient(circle at 30% 30%, #1a1a2e, #0f0f1e);
        }
        
        .controls {
            position: absolute;
            top: 20px;
            left: 20px;
            color: white;
            background: rgba(0, 0, 0, 0.5);
            padding: 20px;
            border-radius: 10px;
            backdrop-filter: blur(10px);
        }
        
        .control-item {
            margin: 10px 0;
        }
        
        label {
            display: inline-block;
            width: 120px;
            font-size: 14px;
        }
        
        input[type="range"] {
            width: 150px;
        }
        
        h3 {
            margin-top: 0;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
        }
    </style>
</head>
<body>
    <div class="controls">
        <h3>🎮 Controls</h3>
        <div class="control-item">
            <label>Speed X:</label>
            <input type="range" id="speedX" min="-5" max="5" value="1" step="0.1">
            <span id="speedXVal">1.0</span>
        </div>
        <div class="control-item">
            <label>Speed Y:</label>
            <input type="range" id="speedY" min="-5" max="5" value="2" step="0.1">
            <span id="speedYVal">2.0</span>
        </div>
        <div class="control-item">
            <label>Sphere Size:</label>
            <input type="range" id="sphereSize" min="20" max="100" value="40" step="5">
            <span id="sphereSizeVal">40</span>
        </div>
    </div>
    
    <canvas id="canvas"></canvas>
    
    <script>
        const canvas = document.getElementById('canvas');
        const ctx = canvas.getContext('2d');
        
        canvas.width = 800;
        canvas.height = 800;
        
        const centerX = canvas.width / 2;
        const centerY = canvas.height / 2;
        
        let angleX = 0;
        let angleY = 0;
        let speedX = 0.01;
        let speedY = 0.02;
        let sphereRadius = 40;
        
        document.getElementById('speedX').addEventListener('input', function(e) {
            speedX = parseFloat(e.target.value) * 0.01;
            document.getElementById('speedXVal').textContent = e.target.value;
        });
        
        document.getElementById('speedY').addEventListener('input', function(e) {
            speedY = parseFloat(e.target.value) * 0.01;
            document.getElementById('speedYVal').textContent = e.target.value;
        });
        
        document.getElementById('sphereSize').addEventListener('input', function(e) {
            sphereRadius = parseInt(e.target.value);
            document.getElementById('sphereSizeVal').textContent = e.target.value;
        });
        
        function project3D(x, y, z) {
            const perspective = 400;
            const scale = perspective / (perspective + z);
            return {
                x: x * scale + centerX,
                y: y * scale + centerY,
                scale: scale
            };
        }
        
        function rotate3D(x, y, z, angleX, angleY) {
            let cosY = Math.cos(angleY);
            let sinY = Math.sin(angleY);
            let xTemp = x * cosY - z * sinY;
            let zTemp = x * sinY + z * cosY;
            x = xTemp;
            z = zTemp;
            
            let cosX = Math.cos(angleX);
            let sinX = Math.sin(angleX);
            let yTemp = y * cosX - z * sinX;
            zTemp = y * sinX + z * cosX;
            y = yTemp;
            z = zTemp;
            
            return { x: x, y: y, z: z };
        }
        
        function drawOctagon(radius, angleX, angleY) {
            const vertices = [];
            
            for (let i = 0; i < 8; i++) {
                const angle = (Math.PI * 2 * i) / 8;
                const x = Math.cos(angle) * radius;
                const y = Math.sin(angle) * radius;
                const z = 0;
                
                const rotated = rotate3D(x, y, z, angleX, angleY);
                const projected = project3D(rotated.x, rotated.y, rotated.z);
                
                vertices.push({ x: projected.x, y: projected.y, z: rotated.z });
            }
            
            ctx.strokeStyle = 'rgba(100, 255, 218, 0.8)';
            ctx.lineWidth = 3;
            ctx.beginPath();
            ctx.moveTo(vertices[0].x, vertices[0].y);
            for (let i = 1; i < vertices.length; i++) {
                ctx.lineTo(vertices[i].x, vertices[i].y);
            }
            ctx.closePath();
            ctx.stroke();
            
            for (let i = 0; i < vertices.length; i++) {
                const v = vertices[i];
                const brightness = (v.z + radius) / (radius * 2);
                ctx.fillStyle = 'rgba(100, 255, 218, ' + (0.3 + brightness * 0.7) + ')';
                ctx.beginPath();
                ctx.arc(v.x, v.y, 6, 0, Math.PI * 2);
                ctx.fill();
            }
        }
        
        function drawSphere(radius, angleX, angleY) {
            const layers = 12;
            const segments = 16;
            
            for (let i = 0; i <= layers; i++) {
                const lat = (Math.PI * i) / layers - Math.PI / 2;
                const layerRadius = Math.cos(lat) * radius;
                const layerY = Math.sin(lat) * radius;
                
                for (let j = 0; j < segments; j++) {
                    const lon = (Math.PI * 2 * j) / segments;
                    const x = Math.cos(lon) * layerRadius;
                    const z = Math.sin(lon) * layerRadius;
                    const y = layerY;
                    
                    const rotated = rotate3D(x, y, z, angleX, angleY);
                    const projected = project3D(rotated.x, rotated.y, rotated.z);
                    
                    if (rotated.z > -radius * 0.3) {
                        const brightness = (rotated.z + radius) / (radius * 2);
                        const size = 2 + brightness * 3;
                        
                        const gradient = ctx.createRadialGradient(
                            projected.x, projected.y, 0,
                            projected.x, projected.y, size
                        );
                        gradient.addColorStop(0, 'rgba(255, 100, 255, ' + (0.4 + brightness * 0.6) + ')');
                        gradient.addColorStop(1, 'rgba(138, 43, 226, ' + (0.2 + brightness * 0.4) + ')');
                        
                        ctx.fillStyle = gradient;
                        ctx.beginPath();
                        ctx.arc(projected.x, projected.y, size, 0, Math.PI * 2);
                        ctx.fill();
                    }
                }
            }
            
            const core = project3D(0, 0, 0);
            const coreGradient = ctx.createRadialGradient(
                core.x, core.y, 0,
                core.x, core.y, sphereRadius * 0.5
            );
            coreGradient.addColorStop(0, 'rgba(255, 150, 255, 0.9)');
            coreGradient.addColorStop(0.5, 'rgba(255, 100, 255, 0.5)');
            coreGradient.addColorStop(1, 'rgba(138, 43, 226, 0.1)');
            
            ctx.fillStyle = coreGradient;
            ctx.beginPath();
            ctx.arc(core.x, core.y, sphereRadius * 0.5, 0, Math.PI * 2);
            ctx.fill();
        }
        
        function animate() {
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            
            angleX += speedX;
            angleY += speedY;
            
            drawOctagon(250, angleX, angleY);
            drawSphere(sphereRadius, angleX * 1.5, angleY * 1.5);
            
            requestAnimationFrame(animate);
        }
        
        animate();
    </script>
</body>
</html>
    `))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	fmt.Println("🚀 3D Octagon Sphere Server started at http://localhost:8080")
	fmt.Println("✨ Открой браузер и наслаждайся 3D эффектом!")
	http.ListenAndServe(":8080", nil)
}
