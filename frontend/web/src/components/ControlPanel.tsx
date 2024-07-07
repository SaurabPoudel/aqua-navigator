import React, { useEffect } from "react";
import axiosInstance from "../axios";

enum Direction {
  RIGHT = "right",
  LEFT = "left",
  UP = "up",
  DOWN = "down",
}

const ControlPanel: React.FC = () => {
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      switch (event.key) {
        case "ArrowRight":
          sendCommand(Direction.RIGHT);
          break;
        case "ArrowLeft":
          sendCommand(Direction.LEFT);
          break;
        case "ArrowUp":
          sendCommand(Direction.UP);
          break;
        case "ArrowDown":
          sendCommand(Direction.DOWN);
          break;
        default:
          break;
      }
    };

    document.addEventListener("keydown", handleKeyDown);

    return () => {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, []);

  const sendCommand = async (cmd: Direction) => {
    try {
      const response = await axiosInstance.post("/control", { command: cmd });
      console.log("Response:", response.data);
    } catch (error) {
      console.error("Error sending command:", error);
    }
  };

  return (
    <div>
      <h2>Control Panel</h2>
    </div>
  );
};

export default ControlPanel;
