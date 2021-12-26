import DateRangePicker from '@wojtekmaj/react-daterange-picker';

function TabCalendar(props) {
    return (
        <DateRangePicker
            onChange={props.getDateRange}
            value={props.defaultDateRange}
            maxDate={new Date()}
        />
    )
}

export default TabCalendar
