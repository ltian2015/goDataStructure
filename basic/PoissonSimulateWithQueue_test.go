package basic

import (
	"fmt"
	"math"
	"math/rand/v2"
	"testing"
)

/*
*

泊松分布\过程(Poisson Distibution\Process)

			 日常生活中，我们会发现很多相互之间没有关联和影响的独立事件，
			 这些独立事件平均发生频率是固定的（经过长期跟踪统计），比如：
							  ● 某医院平均每小时出生3个婴儿
						      ● 某公司平均每10分钟接到1个电话
						      ● 某超市平均每天销售4包xx牌奶粉
						      ● 某网站平均每分钟有2次访问
							  ● 公交车站平均每分钟就有0.2个人到达
							  ● 高速公路上每公里汽车的数量为5
			  但是，给定一个具体的时间段内，发生期望次数事件的概率到底多大呢？
			  这就用到了泊松分布概率计算公式：
			  在给定时间段t，发生n期望次事件的概率记作：P(N(t)=n): ((λ*t)^n)*e^(-λ*t) / n!  这里：
				（1） N(t)是事件数量随着时间变化的函数，N(t)=n表示在时间t内发生了n次事件，P(N(t)=n)表示在时间t内发生n次事件的概率
			    （2）e是欧拉常数，也叫自然常数，是（1+1/x)^x在x趋向无穷大时的一个无限不循环常数，约等于2.71828
				    在golang的math包中预先定义了这个常数：math.E
				（3）λ则是统计出来的事件发生的平均频率（单位时间内的事件次数）
				（4）n是期望发生的事件次数
				（5）t是指定的时间段（注意，时间单位要与λ的时间单位统一）
			  比如，某医院平均每小时出生3个婴儿，那么在接下来的两个小时里，出生0个婴儿的概率为：
							   P（N(2)=0)= (（3个/小时*2小时）^0) *e^(-3个/小时*2小时)/0!
					            =(6^0)*e^(-6)/0!
								=1* 1/e^6 / 1
								≈0.002478752
								≈0.25%
							 在接下来的1个小时里，出生0个婴儿的概率是：
							    P（N（1）=0）= (（3个/小时*1小时）^0) *e^(-3个/小时*1小时)/0!
								 =（3^0)*e^(-3)/0!
								 =1 * 1/e^3 /1
								 ≈0.04979
								 ≈4.979%
					         在接下来的1个小时里，出生1个婴儿的概率是：
							    P（N（1）=1）=(（3个/小时*1小时）^1) *e^(-3个/小时*1小时)/1!
							     =(3^1)*e^(-3)/1!
								 =3 * 1/e^3 /1
								 ≈0.1493
								 ≈14.93%
							 在接下来的1个小时里，出生2个婴儿的概率是：
							     P（N（1）=2）=(（3个/小时*1小时）^2) *e^(-3个/小时*1小时)/2!
								  =(3^2) *e^(-3)/2!
								  =9 * 1/e^3/2
					              ≈0.224
								  ≈22.4%
							 在接下来的1个小时里，出生3个婴儿的概率是：
							 P（N（1）=3））=(（3个/小时*1小时）^3) *e^(-3个/小时*1小时)/3!
							  =(3^3)*e^(-3)/3!
							  =81 * 1 /e^3 /6
							  ≈0.6721
							  ≈67.21%
							  可见，在给定时间段内，越是接近统计出来的平均频率的出生数量的概率就越大。
							  在接下来的1个小时里，出生至少2个婴儿的概率是多少？
							  因为在未来1个小时内出生的婴儿数量可以是0，1，2，3，4...，有无数种可能，所以
							  只能用1-出生0个的概率-出生1个的概率。
							  所以未来1个小时内至少出生两个婴儿的概率为 1-P（N（1）=0）-P（N（1）=1）=1-4.979%-14.93%=80.01%

				对于在在给定的时间段t内到来n个事件的概率 P(N(t)=n): ((λ*t)^n)*e^(-λ*t) / n!
				如果我们已知了λ、n和到达的概率，那么就可以求时间间隔t，这通常用于模拟队列的排队情况，也就是说
				在模拟随机概率情况下，计算后续1个达到事件（n=1）的时间间隔。其思想是：如果下一个事件的发生时间间隔为t，那就意味着
				在t的时间段内没有事件发生，发生了0个事件，其概率为 P(N(t)=0)=((λ*t)^0)*e^(-λ*t) / 0!
				                                                    =e^(-λ*t)
			    故而在t时刻之后发生事件的可能性为P=1-P(N(t)=0)=1-e^(-λ*t)
				如果已知了P，那么时间间隔t就可以计算：
				            因为 P=1-e^(-λ*t)
							     e^(-λ*t)=1-P
						  两边取以e为底的对数（go语言中为Log）
						    -λ*t=Log(1-P)
							t= - Log(1-P)/λ
				这就是在给定概率下，预测下一个事件发生的时间间隔的公式。
				        参见：
							  http://www.360doc.com/content/24/0503/08/48115167_1122206494.shtml
							  https://blog.csdn.net/smartxiong_/article/details/115045758
							  https://preshing.com/20111007/how-to-generate-random-timings-for-a-poisson-process/
                              https://www.johndcook.com/blog/2010/06/14/generating-poisson-random-values/
		  - 泊松概率公式对于服务排队系统研发开发十分有用。
	        我们把每个顾客到来或离开的事件发生时间用Tn表示(n>0)，用a表示到达，用d表示离开，
			用ai表示第i个客户的到达时间（i>=1），用di表示第i个客户的离开时间(i>=1)，
			ci表示第i个客户，line表示排队的情况。
			我们就会形成下面这个随着时间变化的动态排队情况：
		    T0  T1   T2   T3    T4    T5    T6
		    0                                           line: empty
			0   a1                                      line: c1
		    0   a1   a2                                 line: c1, c2
		    0   a1   a2   d1                            line: c2
		    0   a1   a2   d1    a3                      line: c2, c3
		    0   a1   a2   d1    a3    d2                line: c3
		    0   a1   a2   d1    a3    d2    d3          line: empty
		以上表为例，T1时刻，c1刚来，没人排队。
		T2时刻，c2到来，只有c1一个人在排队，这段时间全体客户排队时间为 （T2-T1）*1
		T3时刻，c1离开，此前，有c1和c2两个人排队，这段时间全体客户排队时间为 （T3-T2）*2
		T4时刻，c3刚来，此前，只有c2一个人排队，这段时间全体客户排队时间为 (T4-T3)*1
		T5时刻，c2离开，此前，有c2和c3两人排队，这段时间全体客户排队时间为(T5-T4)*2
		T6时刻，c3离开，此前，只有c3一个人排队，这段时间全体客户排队时间为（T6-T5）*1
		从T6到T1，窗口队伍的全体客户排队时间为：
		    TotalTime=T2-T1）*1+（T3-T2）*2+ (T4-T3)*1+T5-T4)*2+（T6-T5）*1
		平均的排队人数为：TotalTime / (T6-T0)
*/

const (
	arrivalRate           = 0.25 //平均每分钟到达的客户数量
	lowerBoundServiceTime = 0.5  //
	upperBoundServiceTime = 2
	quiteTime             = 480 //480分钟，6个小时
)

func InterArrivalInternal(arrivalRate float64) float64 {
	//建模一个泊松（Poisson）过程并退出
	rn := rand.Float64() //rn在1.0和0.0之间
	//根据泊松分布推导给定概率下，发生一个客户达到的事件的时间间隔。
	//t= - Log(1-P)/λ ,P为给定的概率，λ为平均到达率（单位时间到达的客户数）。
	//Golang math包中，Log是以自然常数e为底的求指数的函数。
	return -math.Log(1.0-rn) / arrivalRate //到达率arrivalRate
}

// 模拟一次随机的服务时间（在给定的服务时间范围内）
// 其实，如果能统计出平均服务率r，那么两个客户离开事件之间的间隔时间（服务时间）也符合泊松分布规律。
// 只不过本模拟没有设定平均服务率，因此，这里只能用一个随机数来表示时间间隔的可能性。
func ServiceTime() float64 {
	rn := rand.Float64()
	return lowerBoundServiceTime + (upperBoundServiceTime-lowerBoundServiceTime)*rn
}

type Customer struct {
	arrivalTime     float64
	serviceDuration float64
}
type Statistics struct {
	waitTimes       []float64 //是队列人数不变时的等待的时间。人数一旦变动（有人离开或加入）就会产生新的等待时间
	queueTime       float64   //累积的排队时间时间
	longestQueue    int       //最大排队人数
	longestWaitTime float64   //最长等待时间
}

// 存储单个客户从到达到离开的等待时间
func (s *Statistics) AddWaitTime(wait float64) {
	s.waitTimes = append(s.waitTimes, wait)
	if wait > s.longestWaitTime {
		s.longestWaitTime = wait
	}
}

// 队伍的总等待时间，保持某个队伍人数状态下的等待时间，当队伍的人数发生变化时才会计算该时间
func (s *Statistics) AddQueueSizeTime(queueSize int, timeAtSize float64) {
	s.queueTime += float64(queueSize) * timeAtSize
}

// 感知队伍的长度变化，如果长度超过已记录的最长值，就更新该最长值。
func (s *Statistics) AddLenth(length int) {
	if length > s.longestQueue {
		s.longestQueue = length
	}
}

func TestDiscreteEventSimulation(t *testing.T) {

	var lastArrivalTime float64 //新客户到达时间
	var departureTime float64   //队头接受服务的客户离开时间，程序中，每个客户到来排队时，该客户所接受的服务时间就已经模拟出来了。
	var lastEventTime float64   // 上次客户到达或者离开的事件所发生的时间

	lastEventTime = 0.0
	//!!!
	//!!!IT技术人员，经常听到的是消息队列或者事件队列，
	//!!!而这里，队列中等待的则是待服务的客户。本质上，队列的存在是为了等待某种服务的“事或物”而
	//!!!存在的一种buffer（缓冲服务对象到来的速度与服务处理速度之间的不平衡），
	//!!!队列中的“事或物”是什么与等待的服务是什么有关，如果等待的服务的处理对象
	//!!!是事件，那么就是事件队列，如果处理的对象的是客户，那就是客户队列，如果处理的是订单，那就是
	//!!!订单队列。如果队列中的“事物”是有自主进入或离开的"主动对象"，比如客户、患者等，改变队列的状态
	//!!!的就是这些主动对象到达和离开的事件（所引起的队列操作）。如果队列中的事物是没有自主权的“被动对象”，
	//!!!比如， 订单或者事件，那么就需要有一个主动对象将其放入或移除队列，这个主动对象就是调用了队列操作的
	//!!!程序所代表的的行为发生者。此时，是这个行为发生者（调用者）与被调用者（队列）之间的协作，完成了
	//!!!被动对象（比如订单或者事件）的状态（在队列中还是出列的）改变。实际上，所有容器对象都是两个对象的合体，
	//!!! 都是一个主动容器管理对象（具有管理行为的对象）和一个被动的容器存储对象（容纳其他对象）的合体，
	//!!! 主动容器管理对象对外隔离和屏蔽了有状态的被动的容器存储对象。相当于仓库管理员对外隔离和屏蔽了仓库的状态。
	//!!! 仓库才是一个有状态的被动业务对象，仓库管理员在自己的意识世界中观察和跟踪仓库的状态，对外暴露仓库的管理行为。
	//!!! 所以，如果设计一个仓库管理微服务，那么，站在这个微服务的外部看，那么这个微服务就是对仓库管理员（组织）的行为建模。
	//!!! 但是，站在这个微服务的内部，就要进一步对仓库管理员意识世界中的仓库这个被动对象进行建模，建模为带有状态的实体。
	//!!! 主动业务对象具有下特点：能够观察和感知事物的状态（状态变或不变是观察与感知的结果）、能够接收信息并做出反映，能够主动发出信息。
	//!!! 对于我们在计算机模拟一台无意识的机械加工零零件的过程，其实这个模拟过程是该机器设计者设计这个机器时所运用的力学与几何计算过程，
	//!!! 并将用不断的计算出来的结果更新及其和零件的状态，实时地用图形展现出来。
	//!!!
	//!!! 回到本话题，这里，服务要处理的是客户，而不是事件，但事件（客户达到或离开队列)确实会改变队列的状态.
	//!!! 如果站在要动态跟踪和管理队列的状态的队列管理员视角看，那么队列的实例，就是一个有ID的实体对象，
	//!!! 该实体以排队等待的服务窗口的启停为生命周期,这个实体的状态（排队的客户）会因为客户到达或离开的事件而改变。
	//!!! 队列管理员会根据客户达到或离开的事件来操作自己所管理的队列。而与队列管理员打交道的其他人是不知道队列内部的状态的。

	line := SliceQueue[Customer]{}
	statistics := Statistics{}

	for { //无限循环相当人不停地在观察被观察的事物，感知和捕捉其变化。这是一种编程模式。
		//表示上一个客户到达后，模拟求取到达时间。
		lastArrivalTime = lastArrivalTime + InterArrivalInternal(arrivalRate)
		//如果最后到达时间超过了工作时间就结束
		if lastArrivalTime > quiteTime {
			break
		}
		if line.Size() == 0 { //如果队伍是空的，前面没人排队, 则客户被直接服务
			lastEventTime = lastArrivalTime //上一个事件时间记作该客户的到达时间
			fmt.Printf("\n 在时间%0.2f 时，有1个客户到达，排队队伍从空无一人变为有1个在等待服务!", lastEventTime)
			ServiceTime := ServiceTime()                       //用变量保留服务持续时间是为了生成接受服务后的离开时间（departureTime）。
			customer := Customer{lastArrivalTime, ServiceTime} //构建一个排队客户
			line.Insert(customer)                              //队列中增加了一个人
			statistics.AddLenth(line.Size())                   //增加排队最长 值的统计
			departureTime = lastArrivalTime + ServiceTime      //计算被服务者离开的时间（此时间之前可能有客户到来，也可能没有客户到来）
		} else { //如果前面有人（队列不空），说明队头的客户在接收服务，程序还不知道是否有人离开，因此分两种情况模拟
			if lastArrivalTime < departureTime { //模拟观察到了新客户到达事件发生在队头接受服务客户离开事件之前的情况				customer := Customer{lastArrivalTime, ServiceTime()}                    //每个客户的服务时间是随机的。
				customer := Customer{lastArrivalTime, ServiceTime()}                    //在客户到达时就假定了该客户接受服务的时长。
				statistics.AddQueueSizeTime(line.Size(), lastArrivalTime-lastEventTime) //队伍在该人数状态所持续时间
				lastEventTime = lastArrivalTime
				line.Insert(customer)
				fmt.Printf("\n 在时间%0.2f 时，有1个客户到达，此时队伍的人数为%d", lastArrivalTime, line.Size())
				statistics.AddLenth(line.Size())
			} else { //模拟观察到了新客户到达时间发生在队头接受服务客户
				statistics.AddQueueSizeTime(line.Size(), departureTime-lastEventTime) //队伍在该人数状态所持续时间
				departuringCustomer := line.Remove()                                  //移除队头服务完毕的客户
				statistics.AddWaitTime(departureTime - departuringCustomer.arrivalTime)
				lastEventTime = departureTime
				fmt.Printf("\n在时间 %0.2f,有客户离开，此时队伍的人数为 %d: ", lastEventTime, line.Size())
				if line.Size() > 0 { //上个客户离开后，计算并更新下一个客户的离开时间， 以便循环观察是有人到达还是离开
					departureTime = lastEventTime + line.First().serviceDuration
				}
			}
		}
	}
	totalWaitTime := 0.0
	//将每个人的等待时间加起来作为总体等待时间
	totalWaitTime = LeftFold(statistics.waitTimes, func(t1, t2 float64) float64 { return t1 + t2 })
	//平均等待时间为总体等待时间除以等待人数。
	averageWaitTime := totalWaitTime / float64(len(statistics.waitTimes))
	fmt.Printf("\n从到达到离开的平均等待时间: %0.2f 分钟", averageWaitTime)
	fmt.Printf("\n等待队伍的平均长度 %0.2f", statistics.queueTime/totalWaitTime)
	fmt.Printf("\n 一天之中，队伍最长为: %d", statistics.longestQueue)
	fmt.Printf("\n一天之中，等待最长的时间为: %0.2f 分钟", statistics.longestWaitTime)
}
