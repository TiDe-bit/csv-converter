import {FC, SetStateAction, useState} from "react";
import { FromToFields } from "./App";

export const AddOptionElement: FC<{initialFrom?: number, initialTo?: number, setFromToPair: (input: FromToFields, remove?: boolean) => void}> = (props) => {
    const {initialFrom, initialTo, setFromToPair} = props
    const [from, setFrom] = useState(initialFrom ?? 0)
    const [to, setTo] = useState(initialTo ?? 0)


    return(
        <div id="input" className="input-box">
            <h2>Add Option: move {from} to {to}</h2>
            <form>
                <input id="from" className="from" onChange={(value) => setFrom(Number(value.target.value) > 0 ? Number(value.target.value) : 0)} autoComplete="off" name="from" type="number" value={from}/>
                <input id="to" className="to" onChange={(value) => setTo(Number(value.target.value) > 0 ? Number(value.target.value):0)} autoComplete="off" name="to" type="number" value={to}/>
                <button className="btn" onClick={() => setFromToPair({From: from, To: to})}>push</button>
                <button className="btn" onClick={() => setFromToPair({From: from, To: to}, true)}>pull</button>
            </form>
        </div>
    );
}