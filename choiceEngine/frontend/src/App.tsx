import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom"
import HomePage from "./home/HomePage";
import GamePage from "./game/GamePage";
import StoryPage from "./story-generation/StoryPage";

function App() {

    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/game" element={<GamePage />} />
                <Route path="/story" element={<StoryPage />}/>
            </Routes>
        </BrowserRouter>

    )
}

export default App
