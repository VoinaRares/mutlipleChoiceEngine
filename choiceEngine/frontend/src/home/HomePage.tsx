import { useNavigate } from "react-router-dom"
import { Initialize } from "../../wailsjs/go/main/App"

// We will want to remove these hard coded values and actually maybe make the user choose a file or let them created a new one
const DIALOG_PATH = "C:\\Users\\voina\\GolandProjects\\learningProject\\choiceEngine\\backend\\dialog.json"
const PLAYER_STATES_PATH = "C:\\Users\\voina\\GolandProjects\\learningProject\\choiceEngine\\backend\\player_states.json"


function HomePage() {
    const navigate = useNavigate()

    async function onClickButton(path :string) {
        try{
            await Initialize(DIALOG_PATH, PLAYER_STATES_PATH)
            navigate(path)
        } catch (err){
            console.log(err)
        }

    }

    return (
        <div>
            <button onClick={() => onClickButton("/game")}>
                Start
            </button>
            <button onClick={() => onClickButton("/story")}>
                Story Generation
            </button>
        </div>
    )
}

export default HomePage