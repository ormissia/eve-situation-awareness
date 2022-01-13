import React, {useEffect} from "react";
import * as echarts from 'echarts'


function Chart(props) {
    useEffect(() => {
        let chartDom = echarts.getInstanceByDom(document.getElementById(props.id));
        if (chartDom == null) {
            chartDom = echarts.init(document.getElementById(props.id));
        }
        chartDom.setOption(props.chartOption)
    }, [props])

    return (
        <div id={props.id} style={{padding: '12px', height: 350}}/>
    );

}

// class Chart extends Component {
//
//     componentDidMount() {
//         this.initChart()
//     }
//
//     componentDidUpdate(prevProps, prevState, snapshot) {
//         //更新图表
//         this.initChart(prevProps);
//     }
//
//     /*生成图表，做了判断，如果不去判断dom有没有生成，
//      每次更新图表都要生成一个dom节点*/
//     initChart(props) {
//         let option = props === undefined ? this.props.chartOption : props.chartOption;
//         // 基于准备好的dom，初始化echarts实例
//         let myChart = echarts.getInstanceByDom(document.getElementById(this.props.id));
//         if( myChart === undefined){
//             myChart = echarts.init(document.getElementById(this.props.id));
//         }
//         // 绘制图表，option设置图表格式及源数据
//         myChart.setOption(option);
//     }
//
//     render() {
//         return (
//             <div id={this.props.id} style={{padding: '12px', height: 350}}/>
//         )
//     }
// }

export default Chart
