import styles from './Sidebar.module.css'

interface MapFieldProps {
    label: string
    values: Record<string, string>
    onChange: (values: Record<string, string>) => void
}

function SidebarMaxField({label, values, onChange}: MapFieldProps) {
    const entries = Object.entries(values)

    function updateEntry(index: number, key: string, value: string) {
        const nextEntries = [...entries]
        nextEntries[index] = [key, value]
        onChange(Object.fromEntries(nextEntries))
    }

    return (
        <fieldset className={styles.field}>
            <legend>{label}</legend>
            <div className={styles.mapField}>
                {entries.map(([key, value], index) => (
                    <div className={styles.mapRow} key={`${label}-${index}`}>

                        <input value={key} placeholder="Key"
                               onChange={(event) =>
                                   updateEntry(index, event.target.value, value)} />

                        <input value={value} placeholder="Value"
                               onChange={(event) =>
                                   updateEntry(index, key, event.target.value)} />

                        <button
                            className={styles.removeButton}
                            onClick={() =>
                                onChange(Object.fromEntries(entries.filter((_, entryIndex) =>
                                    entryIndex !== index)))}
                        >
                            -
                        </button>
                    </div>
                ))}
                <button className={styles.addButton} onClick={() => onChange({...values, '': ''})} type="button">
                    +
                </button>
            </div>
        </fieldset>
    )
}

export default SidebarMaxField
