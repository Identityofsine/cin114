import './styles/screeningheader.scss';

function ScreeningHeader() {
  return (
    <div className="screening-header">
      <img src="/ui/cin114.svg" alt="CIN114" className="screening-header__side" />
      <img src="/ui/cin114_top.svg" alt="CIN114" className="screening-header__top" />
      <span className="screening-header__main">SCREENING</span>
    </div>
  );
}

export default ScreeningHeader
