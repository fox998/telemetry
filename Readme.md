
# Node

Simple command to generate pseudo data for a telemetry sink.

## Command Line Arguments

| Flag | Default Value | Description |
| :--- | :------------ | :---------- |
| **rate** | $1$ | Number of messages per second. |
| **sensor\_name** | `sensor_[1 to 100]` | The name of the sensor to use. |
| **node\_addr** | `localhost:8080` | Address of the telemetry sink. |

-----

# Sink

Aggregates data in memory buffer and flushes it to a file periodically or when the buffer is full.

## Command Line Arguments

| Flag | Default Value | Description |
| :--- | :------------ | :---------- |
| **buffer\_size** | $100$ | Size of the log buffer. |
| **write\_interval** | $100$ | Delta time (in milliseconds) between flush triggers. |
| **log\_file** | `sink.log` | Log file name. |
| **sink\_addr** | `localhost:8080` | Address to listen on. |

-----

# How to Run

1. **Start the Sink:**

    ```bash
    go run ./cmd/sink --buffer_size 50 --write_interval 200 --log_file sink.log
    ```

2. **Start the Node:**

    ```bash
    go run ./cmd/node -rate=5 -name=sensor_1
    ```
