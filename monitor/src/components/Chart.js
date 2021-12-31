import React, {useEffect} from "react";
import * as echarts from 'echarts'


function Chart(props) {
    useEffect(() => {
        let chartDom = echarts.getInstanceByDom(document.getElementById('chart'));
        if (chartDom == null) {
            chartDom = echarts.init(document.getElementById('chart'));
        }
        chartDom.setOption(props.chartOption)
    }, [props])

    return (
        <div id='chart' style={{height: 350}}/>
    );

}

export default Chart
