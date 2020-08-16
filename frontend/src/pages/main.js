import React from 'react'
import { Link } from 'react-router-dom'

const MainPage = () => {
  return (
    <div className='main-page'>
      <h2 className='title'>코로나 사태 속 <strong>예비군 훈련 강행</strong>, 누구를 위한 선택일까요?</h2>
      <section className='item'>
        <h4>
          <strong>국방부, 2020년 9월 1일부터 예비군 소집 훈련 재개 발표</strong>
        </h4>
        <p>
          국방부는 코로나19 확산 방지와 예비군의 안전, 현역부대의 여건 등을 고려하여, 2020년 9월 1일부터 예비군 소집훈련을 하루 일정(개인별 오전, 오후 중 선택)으로 축소 시행하고,
          원격교육은 11월 이후 시험적용하기로 결정하였습니다. (<a href='https://www.yna.co.kr/view/AKR20200729054651504' target='_blank' rel="noopener noreferrer">출처</a>)
        </p>
      </section>

      <section className='item'>
        <h4>
          <strong>예비군 훈련이 5,000만 국민의 건강보다 중요한가요? 정말로요?</strong>
        </h4>
        <p>
          2020년 8월 15일, 정세균 국무총리는 코로나19의 확산세를 "정체절명의 상황"으로 표현하며 서울시와 경기도의 사회적 거리두기를 2단계로 격상한다고 발표했습니다. (<a href='http://news.kbs.co.kr/news/view.do?ncd=4517909&ref=D' target='_blank' rel="noopener noreferrer">출처</a>)
        </p>
        <p>
          코로나19는 아직 끝나지 않았습니다. 코로나19 2020년 8월 13일 추가 확진자는 <strong>103명</strong>, 8월 14일 추가 확진자는 <strong>166</strong>명입니다.(<a href='https://news.daum.net/covid19' target='_blank' rel="noopener noreferrer">출처</a>) 국민들은 아직 불안에 떨고 있습니다.
        </p>
        <p>
          <strong>그런데도 국방부는 코로나19 확산 방지와 예비군의 안전, 현역 부대의 여건을 고려하였다며 9월 1일부터 예비군 훈련을 재개하려고 합니다.</strong>
        </p>
      </section>

      <section className='item'>
        <h4>
          <strong>예비군 훈련은 취소되어야 합니다.</strong>
        </h4>
        <p> 
          대한민국 예비군 편성 인원은 약 300만 명입니다. 정부가 나서서 100인 이상의 모임/행사도, 공공 다중 시설 이용도 금지하고 있는 마당에 예비군 훈련이라니요? 예비군 훈련은 취소되어야 합니다.
        </p>
        <p>
          21대 국회 국방위원회 및 보건복지위원회 의원님들께 메일을 보내 의견을 전달해주세요.
        </p>
      </section>
      <div className='mail-button-wrap'>
        <Link to='/petition'><button className='mail-button'>메일 보내러 가기</button></Link>
      </div>
    </div>
  )
}

export default MainPage
