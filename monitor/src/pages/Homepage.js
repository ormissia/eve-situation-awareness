import React, {Component} from "react";
import {Layout, Button, Menu,Typography} from 'antd';

import {API_GET_SOLAR_SYSTEM_KILL} from "../utils/Api";

const {Header, Footer, Content} = Layout;
const { Title } = Typography;

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

    getTabMenuTimeType = (timeType, startTS, endTS, solarSystemName) => {
        API_GET_SOLAR_SYSTEM_KILL(timeType).get(
            '/solar_system_kill?time_type=' + timeType
            + '&start_time_stamp=' + startTS
            + '&end_time_stamp=' + endTS
            + '&page_size='
            + '&solar_system_name=' + solarSystemName)
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
            <Layout style={{background: 'white'}}>
                <Header style={{position: 'fixed', zIndex: 1, width: '100%', background: 'whitesmoke'}}>
                    <Menu mode="horizontal" defaultSelectedKeys={['2']}>
                        {new Array(6).fill(null).map((_, index) => {
                            const key = index + 1;
                            return <Menu.Item key={key}>{`nav ${key}`}</Menu.Item>;
                        })}
                    </Menu>
                </Header>
                <Content className="site-layout" style={{padding: '0 50px', marginTop: 64}}>
                    <div style={{heigh: '10000px', color: 'red'}}>
                        <Button type="primary">Button</Button>
                    </div>
                </Content>
                <Footer style={{background: 'whitesmoke'}}>
                    foooooot
                </Footer>
            </Layout>
        )
    }
}


export default Homepage
