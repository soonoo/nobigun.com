import React, { useEffect, useState, useCallback } from 'react'
import axios from 'axios'
import { useHistory } from "react-router-dom";

const HOST = process.env.NODE_ENV === 'development'
  ? 'http://localhost:1323'
  : 'https://api.nobigun.com'

const PetitionPage = () => {
  const history = useHistory()
  const [recipients, setRecipients] = useState([])
  const [isLoading, setIsLoading] = useState(false)
  const [formData, updateFormData] = React.useState({
    to: '0',
    senderName: '',
    senderAddress: '',
    content: '',
  });

  const handleChange = useCallback((e) => {
    updateFormData({
      ...formData,
      [e.target.name]: e.target.value.trim()
    });
  }, [formData]);

  const onSubmit = useCallback((e) => {
    e.preventDefault()
    if (isLoading) {
      return
    }
    if (formData.senderName === '') {
      alert('성함을 입력해주세요.')
      return
    }

    console.log(formData)

    setIsLoading(true)
    axios.post(`${HOST}/petitions`, {
      ...formData,
      to: formData.to === '0' ? null : formData.to,
    })
      .then(()=> {
        alert('이메일을 전송했습니다. 감사합니다!')
        history.push('/')
      })
      .finally(() => {
        setIsLoading(false)
      })
  }, [formData, isLoading, history])

  useEffect(() => {
    axios.get(`${HOST}/recipients`)
      .then((r) => {
        setRecipients(r.data || [])
      })
  }, [])

  return (
    <div className='petition-page'>
      <h2 className='title'>21대 국회 <strong>국방위원회</strong> 및 <strong>보건복지위원회</strong> 소속 의원들에게 대신 의견을 전달해드릴게요.</h2>
      <p>
        진정성 있는 의견을 전달을 위해 장난스런 메세지는 자제 부탁드립니다.
      </p>
      <form onSubmit={onSubmit}>
        <div className='row'>
          <div className='twelve columns'>
            <label htmlFor='exampleRecipientInput'>
              누구에게 보낼까요?
              <select className='u-full-width' name='to' onChange={handleChange}>
                <option value='0'>무작위로 보내기</option>
                {recipients.map((r) => {
                  const { id, name, group, email } = r
                  if (!email) return null
                  return <option value={id} key={id}>{name} ({group})</option>
                })}
              </select>
            </label>
          </div>
        <div className='row'>
          <div className='twelve columns'>
            <label>
              성함 (필수, 이메일 발송을 위해서만 사용됩니다.)
              <input className='u-full-width' type='text' placeholder='홍길동' name='senderName' onChange={handleChange} />
            </label>
          </div>
            </div>
        <div className='row'>
          <div className='twelve columns'>
            <label>
              답장 받을 이메일 (선택, 입력하시면 의원님들이 답장을 주실지도 몰라요.)
              <input className='u-full-width' type='email' placeholder='name@gmaih.com' name='senderAddress' onChange={handleChange} />
            </label>
          </div>
            </div>
        <div className='row'>
          <div className='twelve columns'>
            <label>
              전달할 내용 (선택, 입력하시지 않으면 기본 템플릿 메일을 전송합니다.)
              <textarea rows='10' className='u-full-width' placeholder='정책 담당자에게 전달할 메세지를 입력해주세요.' name='content' onChange={handleChange}></textarea>
            </label>
          </div>
        </div>
          </div>
        <div className='row'>
          <div className='twelve columns'>
        <div className='mail-button-wrap'>
          <button className='mail-button u-full-width' type='submit'>메일 보내기</button>
        </div>
          </div>
            </div>
      </form>
    </div>
  )
}

export default PetitionPage
