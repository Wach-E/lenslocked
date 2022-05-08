# Rectify issue with ports
netstat -ano | findstr :thePortOfInterest
taskkill /PID portPID /F 