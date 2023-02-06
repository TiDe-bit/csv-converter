import {ReactNode, useState} from 'react';
import './App.css';
import {SpreadColumns, AddOption, Convert, RemoveOption} from "../wailsjs/go/app/App";
import {addOption} from "./addOption";

function App() {
    const [OptionThingies, setEndArray] = useState<ReactNode[]>([]);


    function convert() {
        Convert()
    }

    function  spreadColumns (): ReactNode {
        let columns: string[] = []
        SpreadColumns().then((c) => columns = c)
        if (columns.length === 0) {
            return <li>nothing here</li>
        }
        return columns.map((column,  index) => (<li>{index} - {column}</li>))
    }

    return (
        <div id="App">
            <ul id="rowNames" className="RowNames">{spreadColumns()}</ul>
            <div id="input" className="input-box">
                <button className="btn" onClick={() => setEndArray([...OptionThingies,addOption()])}>AddOption</button>
                <button className="btn" onClick={convert}>Convert</button>
            </div>
            {OptionThingies}
        </div>
    )
}

export default App
