import {Row, Col} from 'antd';
import {Component} from "react";

import Chart from "./Chart";

import {API_GET_SOLAR_SYSTEM_KILL} from "../utils/Api";
import {BigNumFormat} from "../utils/Utils";

class HomepageContent extends Component {
    state = {
        killNumValueChartOption: {
            color: ['#B0C4DE', '#191970'],
            title: {
                text: '击杀趋势',
                left: 'center',
                bottom: '0',
            },
            tooltip: {},
            xAxis: {
                data: []
            },
            yAxis: [
                {
                    splitLine: {show: false},
                    type: 'value',
                    name: 'kill quantity',
                    splitNumber: 5, //设置坐标轴的分割段数
                },
                {
                    splitLine: {show: false},
                    type: 'value',
                    name: 'kill value',
                    splitNumber: 5, //设置坐标轴的分割段数
                    axisLabel: {
                        formatter: BigNumFormat,
                    },
                }
            ],
            series: [
                {
                    name: 'kill quantity',
                    type: 'bar',
                    yAxisIndex: 0,
                    data: []
                },
                {
                    name: 'kill value',
                    type: 'line',
                    smooth: true,
                    yAxisIndex: 1,
                    data: []
                }
            ]
        },
        killValueTopNChartOption: {
            color: ['#B0C4DE', '#191970'],
            title: {
                text: '击杀价值TopN',
                left: 'center',
                bottom: '0',
            },
            tooltip: {},
            xAxis: {
                data: []
            },
            yAxis: [
                {
                    splitLine: {show: false},
                    type: 'value',
                    name: 'kill value',
                    splitNumber: 5, //设置坐标轴的分割段数
                    axisLabel: {
                        formatter: BigNumFormat,
                    },
                },
            ],
            series: [
                {
                    name: 'kill value',
                    type: 'bar',
                    yAxisIndex: 0,
                    data: []
                },
            ]
        },
        killQuantityTopNChartOption: {
            color: ['#B0C4DE', '#191970'],
            title: {
                text: '击杀数量TopN',
                left: 'center',
                bottom: '0',
            },
            tooltip: {},
            xAxis: {
                data: []
            },
            yAxis: [
                {
                    splitLine: {show: false},
                    type: 'value',
                    name: 'kill quantity',
                    splitNumber: 5, //设置坐标轴的分割段数
                },
            ],
            series: [
                {
                    name: 'kill value',
                    type: 'bar',
                    yAxisIndex: 0,
                    data: []
                },
            ]
        },
    }

    //子组件完成挂载时, 将子组件的方法 this 作为参数传到父组件的函数中
    componentDidMount() {
        // 在子组件中调用父组件的方法,并把当前的实例传进去
        this.props.onChildEvent(this)
        this.loadData()
    }

    // 子组件的方法, 在父组件中触发
    loadData = () => {
        this.getKillNumValueChartOption(this.props.conditions.group, this.props.conditions.startTime, this.props.conditions.endTime, this.props.conditions.solarSystem)
        this.getKillValueTopNChartOption(this.props.conditions.group, this.props.conditions.startTime, this.props.conditions.endTime, 'kill_quantity')
        this.getKillValueTopNChartOption(this.props.conditions.group, this.props.conditions.startTime, this.props.conditions.endTime, 'kill_value')
    }

    getKillNumValueChartOption = (timeType, startTime, endTime, solarSystemName) => {
        API_GET_SOLAR_SYSTEM_KILL(timeType).get(
            '/solar_system_kill?time_type=' + timeType
            + '&start_time_stamp=' + startTime
            + '&end_time_stamp=' + endTime
            + '&page_size='
            + '&solar_system_name=' + solarSystemName)
            .then(response => {
                console.log(response.data.data)
                const initOption = {
                    xAxis: {
                        data: response.data.data.X,
                    },
                    series: [
                        {
                            name: 'kill quantity',
                            type: 'bar',
                            yAxisIndex: 0,
                            data: response.data.data.Y.kill_quantity
                        },
                        {
                            name: 'kill value',
                            type: 'line',
                            smooth: true,
                            yAxisIndex: 1,
                            data: response.data.data.Y.kill_value
                        }
                    ]
                }
                const newOption = Object.assign(this.state.killNumValueChartOption, initOption)
                this.setState({killNumValueChartOption: newOption})
            })
    }

    getKillValueTopNChartOption = (timeType, startTime, endTime, order_field) => {
        API_GET_SOLAR_SYSTEM_KILL(timeType).get(
            '/solar_system_kill_order?time_type=' + timeType
            + '&start_time_stamp=' + startTime
            + '&end_time_stamp=' + endTime
            + '&order_field=' + order_field
            + '&order_type=desc')
            .then(response => {
                console.log(response.data.data)
                const initOption = {
                    xAxis: {
                        data: response.data.data.X,
                        axisLabel: {
                            interval: 0,
                            rotate: 30,
                        },
                    },
                    series: [
                        {
                            name: order_field,
                            type: 'bar',
                            yAxisIndex: 0,
                            data: (order_field === 'kill_quantity') ? response.data.data.Y.kill_quantity : response.data.data.Y.kill_value,
                        }
                    ]
                }
                if (order_field === 'kill_quantity') {
                    const newOption = Object.assign(this.state.killQuantityTopNChartOption, initOption)
                    this.setState({killQuantityTopNChartOption: newOption})
                } else if (order_field === 'kill_value') {
                    const newOption = Object.assign(this.state.killValueTopNChartOption, initOption)
                    this.setState({killValueTopNChartOption: newOption})
                }
            })
    }

    render() {
        return (
            <div>
                <Row gutter={[16, 24]}>
                    <Col className="gutter-row" span={18}>
                        <Row gutter={[16, 24]}>
                            <Col className="gutter-row" span={24}>
                                <Chart
                                    id={'killNumValueChartOption'}
                                    chartOption={this.state.killNumValueChartOption}
                                />
                            </Col>
                            <Col className="gutter-row" span={12}>
                                <Chart
                                    id={'killQuantityTopNChartOption'}
                                    chartOption={this.state.killQuantityTopNChartOption}
                                />
                            </Col>
                            <Col className="gutter-row" span={12}>
                                <Chart
                                    id={'killValueTopNChartOption'}
                                    chartOption={this.state.killValueTopNChartOption}
                                />
                            </Col>
                        </Row>
                    </Col>
                    <Col className="gutter-row" span={6}>
                        <Row gutter={[16, 24]}>
                            <Col className="gutter-row" span={24}>
                                <div style={{background: '#0092ff', padding: '8px 0'}}>危险星系</div>
                            </Col>
                            <Col className="gutter-row" span={24}>
                                <div style={{background: '#0092ff', padding: '8px 0'}}>危险角色</div>
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </div>
        );
    }
}

export default HomepageContent
