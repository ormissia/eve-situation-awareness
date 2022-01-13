import {DatePicker, Input, Row, Col, Select} from 'antd';
import moment from 'moment';

const {RangePicker} = DatePicker;
const {Option} = Select;

function HomepageHeader(props) {

    function disabledDate(current) {
        // Can not select days after today
        return current > moment().endOf('day');
    }

    function selectHandleChange(value) {
        props.initConditions.group = value
        props.callback(props.initConditions)
    }

    function rpOnChange(dates) {
        if (dates != null) {
            props.initConditions.startTime = dates[0].valueOf()
            props.initConditions.endTime = dates[1].valueOf()
            props.callback(props.initConditions)
        }
    }

    return (
        <div>
            <Row gutter={[16, 24]}>
                <Col className="gutter-row" span={5} offset={11}>
                    <RangePicker
                        disabledDate={disabledDate}
                        onChange={rpOnChange}
                        // value={[moment(props.initConditions.startTime), moment(props.initConditions.endTime)]}
                        defaultValue={[moment(props.initConditions.startTime), moment(props.initConditions.endTime)]}
                        ranges={{
                            'Last 24 Hours': [moment().subtract(1, 'days'), moment()],
                            'Last Week': [moment().startOf('day').add(1, 'days').subtract(1, 'weeks'), moment()],
                            'Last Month': [moment().startOf('day').add(1, 'days').subtract(1, 'months'), moment()],
                            'Last Year': [moment().startOf('day').add(1, 'days').subtract(1, 'years'), moment()],
                            'This Month': [moment().startOf('month'), moment()],
                            'This Year': [moment().startOf('year'), moment()],
                        }}
                    />
                </Col>
                <Col className="gutter-row" span={3}>
                    Group&nbsp;&nbsp;
                    <Select defaultValue="hour" style={{width: 90}} onChange={selectHandleChange}>
                        <Option value="hour">Hour</Option>
                        <Option value="day">Day</Option>
                        <Option value="month">Month</Option>
                        <Option value="year">Year</Option>
                    </Select>
                </Col>
                <Col className="gutter-row" span={4}>
                    <Input placeholder="input search text"/>
                </Col>
            </Row>
        </div>
    );
}

export default HomepageHeader
