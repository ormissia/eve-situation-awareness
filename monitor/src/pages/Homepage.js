import {
    Header,
    Container,
    Segment,
} from 'semantic-ui-react'
import React, {Component} from "react";

import TabMenu from "../components/TabMenu";
import Chart from "../components/Chart";
import API_GET_SOLAR_SYSTEM_KILL from "../utils/Api";

const style = {
    h1: {
        padding: '2em',
    },
    h2: {
        margin: '4em 0em 2em',
    },
    h3: {
        marginTop: '2em',
        padding: '2em 0em',
    },
    last: {
        marginBottom: '300px',
    },
}

class Homepage extends Component {

    state = {
        chartOption: {
            color: ['#B0C4DE','#191970'],
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
        }
    }

    getTabMenuTimeType = (timeType, startTS, endTS) => {
        API_GET_SOLAR_SYSTEM_KILL(timeType).get('/solar_system_kill?time_type=' + timeType + '&start_time_stamp=' + startTS + '&end_time_stamp=' + endTS + '&page_size=')
            .then(response => {
                const initOption = {
                    xAxis: {
                        data: response.data.data.X
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
                const newOption = Object.assign(this.state.chartOption, initOption)
                this.setState({chartOption: newOption})
            })
    }

    render() {
        return (
            <div>
                <Container>
                    <TabMenu getTabMenuTimeType={this.getTabMenuTimeType}/>

                    <Header as='h1' content='EVE Situation Awareness' style={style.h1} textAlign='center'/>

                    <Segment>
                        <Chart chartOption={this.state.chartOption}/>
                    </Segment>
                </Container>
            </div>
        )
    }
}


export default Homepage
