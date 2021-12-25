import * as echarts from 'echarts'
import React, {useEffect} from "react";


function Chart(props) {

    // componentDidMount() {
    //     // 基于准备好的dom，初始化echarts实例
    //     let myChart = echarts.init(document.getElementById('chart'));
    //
    //
    //     // 绘制图表
    //     myChart.setOption(this.props.chartOption)
    // }



    useEffect(() => {
        let chartDom = echarts.init(document.getElementById('chart'));
        // var chartDom = echarts.getInstanceById('chart');
        // if (chartDom === 'undefined'){
        //     chartDom = echarts.init(document.getElementById('chart'));
        // }
        chartDom.setOption(props.chartOption)
    }, [props])
    return (
        // <div id="chart" style={{height: 350}}/>
        <div id='chart' style={{height: 350}}/>
    );

}

export default Chart
