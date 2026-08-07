import styles from './Sidebar.module.css'

interface ArrayFieldProps {
    label: string
    values: string[]
    onChange: (values: string[]) => void
}

function SidebarArrayField({label, values, onChange}: ArrayFieldProps) {
    return (
        <fieldset className={styles.field}>
            <legend>{label}</legend>
            <div className={styles.arrayField}>
                {values.map((value, index) => (
                    <div className={styles.arrayRow} key={`${label}-${index}`}>
                        <input
                            value={value}
                            onChange={(event) => {
                                const nextValues = [...values]
                                nextValues[index] = event.target.value
                                onChange(nextValues)
                            }}
                        />
                        <button
                            className={styles.removeButton}
                            onClick={() =>
                                onChange(values.filter((_, valueIndex) => valueIndex !== index))}
                        >
                            -
                        </button>
                    </div>
                ))}
                <button className={styles.addButton} onClick={() => onChange([...values, ''])} type="button">
                    +
                </button>
            </div>
        </fieldset>
    )
}

export default SidebarArrayField
