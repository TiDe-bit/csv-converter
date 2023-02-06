import {ReactNode, useState} from "react";
import {AddOption, RemoveOption} from "../wailsjs/go/app/App";

export function addOption ():ReactNode {
    const [from, setFrom] = useState(0)
    const [to, setTo] = useState(0)
    console.log(from,to)
    async function pushOption(event: any): Promise<void> {
        event.preventDefault()
        await AddOption(event.target.from.value, event.target.to.value, false)
    }
    async function pullOption(): Promise<void> {
        await RemoveOption(from, to)
    }

    return(
        <div id="input" className="input-box">
            <h2>Add Option</h2>
            <input id="from" className="from" ref={(c) => setFrom(Number(c?.value))} autoComplete="off" name="from" type="number"/>
            <input id="to" className="to" ref={(c) => setTo(Number(c?.value))} autoComplete="off" name="to" type="number"/>
            <button className="btn" onClick={pushOption}>push</button>
            <button className="btn" onClick={pullOption}>pull</button>
        </div>)
}