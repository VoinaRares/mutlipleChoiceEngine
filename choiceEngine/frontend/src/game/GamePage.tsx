import { useState, useEffect } from "react"
import { GetCurrent, GetOptions, ChooseOption } from "../../wailsjs/go/main/App"
import {traversal} from "../../wailsjs/go/models";
import Option = traversal.Option;

function GamePage() {
    const [current, setCurrent] = useState<Option | null>(null)
    const [options, setOptions] = useState<Option[]>([])

    async function refresh() {
        const cur = await GetCurrent()
        const opts = await GetOptions(cur.Id)
        setCurrent(cur)
        setOptions(opts)
    }

    useEffect(() => {
        refresh()
    }, [])

    async function onOptionClick(optionId: string) {
        await ChooseOption(optionId)
        await refresh()
    }

    return (
        <div>
            <p>{current?.Text}</p>
            {options.map(opt => (
                <button key={opt.Id} onClick={() => onOptionClick(opt.Id)}>
                    {opt.PreviewText}
                </button>
            ))}
        </div>
    )
}

export default GamePage