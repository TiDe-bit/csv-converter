import {ReactNode, useState} from 'react';
import './App.css';
import {SpreadColumns,  Convert, GetSavedOptions} from "../wailsjs/go/app/App";
import { AddOptionElement } from './addOption';

function App() {
    const [OptionThingies, setEndArray] = useState<ReactNode[]>([]);

    GetSavedOptions().then((savedOptions) => {
        if (savedOptions.length == 0) {
            return
        }
        setEndArray([...OptionThingies, savedOptions.map((option) => <AddOptionElement initialFrom={option.form} initialTo={option.to} />)])
    }) 

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
                <button className="btn" onClick={() => setEndArray([...OptionThingies, <AddOptionElement />])}>AddOption</button>
                <button className="btn" onClick={convert}>Convert</button>
            </div>
            {OptionThingies}
        </div>
    )
}

export default App
