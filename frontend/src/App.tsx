import {ReactElement, ReactNode, useEffect, useState} from 'react';
import './App.css';
import {SpreadColumns, GetSavedOptions, SaveOptions, Convert} from "../wailsjs/go/app/App";
import { AddOptionElement } from './addOption';


export interface FromToFields {
    From: number
    To: number
}

const allOptions = new Map<string, FromToFields>


function App() {
    const [elements, setElements] = useState<ReactElement[]>([])
    const [optionCount, setOptionCount] = useState(0)

    const setFromToPair = (fromTo: FromToFields, remove?: boolean) => {
        console.log("add", fromTo, remove)
        if (remove) {
            allOptions.delete(fromTo.To.toString())
            return
        }
        allOptions.set(fromTo.To.toString(), fromTo)
        setOptionCount(allOptions.size)
        console.log(optionCount, allOptions)
    }

    useEffect(() => {
        let savedOptions: FromToFields[] = [];
        GetSavedOptions().then((value) => {
            savedOptions = value
        });
        savedOptions.forEach((option) => allOptions.set(option.To.toString(), option))
        setOptionCount(allOptions.size)
    })

    useEffect(() => {
        const list: FromToFields[] = []
        console.log("fuck", list, allOptions)
        allOptions.forEach((option: FromToFields) => list.push(option))
        setElements(list.map((fields) => (<>
            <AddOptionElement setFromToPair={setFromToPair} initialFrom={fields.From} initialTo={fields.To} key={fields.To} />
        </>)))
    }, [optionCount])

    function  spreadColumns (): ReactNode {
        let columns: string[] = []
        SpreadColumns().then((c) => columns = c)
        if (columns.length === 0) {
            return <li>nothing here</li>
        }
        return columns.map((column,  index) => (<li>{index} - {column}</li>))
    }

    const convertAndSave = (): void => {
        SaveOptions(allOptions.values as any);
        Convert()
    }

    return (
        <div id="App">
            <ul id="rowNames" className="RowNames">{spreadColumns()}</ul>
            <div id="input" className="input-box">
                <button className="btn" onClick={() => {
                    allOptions.set("0", {From: 0, To: 0})
                    setOptionCount(allOptions.size)
                    console.log(allOptions, elements)
                }}>Add Option</button>
                <button className="btn" onClick={convertAndSave}>Convert</button>
            </div>
            {elements}
        </div>
    )
}

export default App
