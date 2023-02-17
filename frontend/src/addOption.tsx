import {FC, useState} from "react";
import {AddOption, RemoveOption} from "../wailsjs/go/app/App";

export const AddOptionElement: FC<{initialFrom?: number, initialTo?: number}> = (props) => {
    const [from, setFrom] = useState(props.initialFrom ?? 0)
    const [to, setTo] = useState(props.initialTo ?? 0)

    return(
        <div id="input" className="input-box">
            <h2>Add Option</h2>
            <form>
                <input id="from" className="from" onChange={(value) => setFrom(Number(value.target.value))} autoComplete="off" name="from" type="number" value={from}/>
                <input id="to" className="to" onChange={(value) => setTo(Number(value.target.value))} autoComplete="off" name="to" type="number" value={to}/>
                <button className="btn" onClick={() => AddOption(from, to, false)}>push</button>
                <button className="btn" onClick={() => RemoveOption(from, to)}>pull</button>
            </form>
        </div>
    );
}