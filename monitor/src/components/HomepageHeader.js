import {DatePicker, Input, Row, Col,Menu, Dropdown} from 'antd';
import moment from 'moment';

const {RangePicker} = DatePicker;

function HomepageHeader(props) {
    function disabledDate(current) {
        // Can not select days after today
        return current > moment().endOf('day');
    }


    function rpOnChange(dates, dateStrings) {
        if (dates.length === 2) {
            const startTs = dates[0].valueOf()
            const endTs = dates[1].valueOf()
            console.log(startTs)
            console.log(endTs)
        }
    }

    return (
        <div>
            <Row gutter={[16, 24]}>
                <Col className="gutter-row" span={5} offset={14}>
                    <RangePicker
                        disabledDate={disabledDate}
                        onChange={rpOnChange}
                        ranges={{
                            'Last 24 Hours': [moment().subtract(1, 'days'), moment()],
                            'Last Week': [moment().subtract(1, 'weeks'), moment()],
                            'Last Month': [moment().subtract(1, 'months'), moment()],
                            'Last Year': [moment().subtract(1, 'years'), moment()],
                            'This Month': [moment().startOf('month'), moment()],
                            'This Year': [moment().startOf('year'), moment()],
                        }}
                    />
                </Col>
                <Col className="gutter-row" span={4}>
                    <Input placeholder="input search text"/>
                </Col>
            </Row>
        </div>
    );
}

export default HomepageHeader
