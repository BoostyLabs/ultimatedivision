(self.webpackChunkcryptofotball=self.webpackChunkcryptofotball||[]).push([[465],{8995:(e,i,a)=>{"use strict";e.exports=a.p+"images/eye5367e.svg"},4516:(e,i,a)=>{"use strict";e.exports=a.p+"images/futd1cdf.svg"},2003:(e,i,a)=>{"use strict";e.exports=a.p+"images/parametrese33b3.svg"},4867:(e,i,a)=>{"use strict";e.exports=a.p+"images/rectanglefd933.svg"},98:(e,i,a)=>{"use strict";e.exports=a.p+"images/searche524d.svg"},5770:(e,i,a)=>{"use strict";e.exports=a.p+"images/star71ea1.svg"},8754:(e,i,a)=>{"use strict";e.exports=a.p+"images/nextda86e.svg"},110:(e,i,a)=>{"use strict";e.exports=a.p+"images/previous1f4cf.svg"},3382:(e,i,a)=>{"use strict";a.d(i,{R:()=>m});var s=a(655),t=a(5893),r=a(7294),n=a(4867),l=a(98),c=a(5770),o=a(4516),d=a(8995),p=a(2003),u=function(e){var i=e.props,a=i.label,n=i.image,l=(0,r.useState)(!1),c=l[0],o=l[1];return(0,t.jsxs)("div",(0,s.pi)({className:"filter-item",onClick:function(){return o((function(e){return!e}))}},{children:[(0,t.jsx)("span",(0,s.pi)({className:"filter-item__title"},{children:a}),void 0),(0,t.jsx)("img",{className:"filter-item__picture",src:n,alt:n&&"filter icon"},void 0),(0,t.jsx)("div",{className:"filter-item__dropdown",style:{display:c?"block":"none"}},void 0)]}),void 0)},m=function(e){var i=e.title,a=(0,r.useState)(""),m=a[0],v=a[1],_=function(e,i){this.label=e,this.image=i},g=[new _("Version",n),new _("Positions",n),new _("Nations",n),new _("Leagues",n),new _("WRF",n),new _("Stats",n),new _("",c),new _("PS",o),new _("T&S",n),new _("",d),new _("",c),new _("RPP",n),new _("",p),new _("Misc",n)];return(0,t.jsxs)("section",(0,s.pi)({className:"filter-field"},{children:[(0,t.jsx)("h1",(0,s.pi)({className:"filter-field__title"},{children:i}),void 0),(0,t.jsx)("div",(0,s.pi)({className:"filter-field__wrapper"},{children:(0,t.jsxs)("ul",(0,s.pi)({className:"filter-field__list"},{children:[(0,t.jsxs)("li",(0,s.pi)({className:"filter-field__list__item"},{children:[(0,t.jsx)("img",{src:l,alt:"Filter Icon",className:"filter-field__list__item__picture"},void 0),(0,t.jsx)("input",{value:m,placeholder:"Search",className:"filter-field__list__item__search",onChange:function(e){v(e.target.value)}},void 0)]}),void 0),g.map((function(e,i){return(0,t.jsx)("li",(0,s.pi)({className:"filter-field__list__item"},{children:(0,t.jsx)(u,{props:e},void 0)}),i)}))]}),void 0)}),void 0)]}),void 0)}},7928:(e,i,a)=>{"use strict";a.d(i,{D:()=>o});var s=a(655),t=a(5893),r=a(7294),n=a(8754),l=a(110),c=function(e){var i=e.blockPages,a=e.onPageChange;return(0,t.jsx)("ul",(0,s.pi)({className:"ultimatedivision-paginator__pages"},{children:i.map((function(e,i){return(0,t.jsx)("li",(0,s.pi)({className:"ultimatedivision-paginator__pages__item",onClick:function(){return a("change page",e)}},{children:e}),i)}))}),void 0)},o=function(e){var i=e.itemCount,a=(0,r.useState)(1),o=a[0],d=a[1],p=(0,r.useState)([]),u=p[0],m=p[1],v=(0,r.useState)([]),_=v[0],g=v[1],f=(0,r.useState)([]),x=f[0],h=f[1];(0,r.useEffect)((function(){y()}),[o]);for(var j=[],N=1;N<=Math.ceil(i/8);N++)j.push(N);var k=_.length<=5&&j.length>10,C=!!_.length,w=o<5,b=j.length-o<4,P=j.length<=10,y=function(){if(j.length)return P?(m(j.slice()),g([]),void h([])):void function(){if(!P)w?(m(j.slice(0,5)),g([]),h(j.slice(-1))):w||b?b&&(m(j.slice(0,1)),g([]),h(j.slice(-5))):(m(j.slice(0,1)),g(j.slice(o+-3,o+2)),h(j.slice(-1)))}()},S=function(e,i){switch(void 0===i&&(i=o),e){case"next page":return i<j.length&&d(i+1),void y();case"previous page":return i>1&&d(i-1),void y();case"change page":return d(i),void y();default:y()}};return(0,t.jsx)("section",(0,s.pi)({className:"ultimatedivision-paginator"},{children:(0,t.jsxs)("div",(0,s.pi)({className:"ultimatedivision-paginator__wrapper"},{children:[(0,t.jsxs)("a",(0,s.pi)({className:"ultimatedivision-paginator__previous",onClick:function(){return S("previous page")}},{children:[(0,t.jsx)("img",{className:"ultimatedivision-paginator__previous__arrow",src:l,alt:"Previous page"},void 0),(0,t.jsx)("p",(0,s.pi)({className:"ultimatedivision-paginator__previous__title"},{children:"Previous page"}),void 0)]}),void 0),(0,t.jsx)(c,{blockPages:u,onPageChange:S},void 0),k&&(0,t.jsx)("span",(0,s.pi)({className:"ultimatedivision-paginator__pages__dots"},{children:"..."}),void 0),(0,t.jsx)(c,{blockPages:_,onPageChange:S},void 0),C&&(0,t.jsx)("span",(0,s.pi)({className:"ultimatedivision-paginator__pages__dots"},{children:"..."}),void 0),(0,t.jsx)(c,{blockPages:x,onPageChange:S},void 0),(0,t.jsxs)("a",(0,s.pi)({className:"ultimatedivision-paginator__next",onClick:function(){return S("next page")}},{children:[(0,t.jsx)("p",(0,s.pi)({className:"ultimatedivision-paginator__next__title"},{children:"Next page"}),void 0),(0,t.jsx)("img",{className:"ultimatedivision-paginator__next__arrow",src:n,alt:"Next page"},void 0)]}),void 0)]}),void 0)}),void 0)}},6317:(e,i,a)=>{"use strict";a.d(i,{s:()=>l});var s=a(655),t=a(5893),r=a(3727),n=a(3356),l=function(e){var i=e.card,a=e.parentClassName;return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)("img",{className:a+"__background-type",src:i.mainInfo.backgroundType,alt:"background img",draggable:!1},void 0),(0,t.jsx)("img",{className:a+"__face-picture",src:i.mainInfo.playerFace,alt:"Player face",draggable:!1},void 0),(0,t.jsx)(r.rU,(0,s.pi)({to:{pathname:n.FM.FootballerCard.path,state:{card:i}}},{children:(0,t.jsx)("span",(0,s.pi)({className:a+"__name"},{children:i.mainInfo.lastName}),void 0)}),void 0),(0,t.jsx)("ul",(0,s.pi)({className:a+"__list"},{children:i.stats.map((function(e,i){return(0,t.jsx)("li",(0,s.pi)({className:a+"__list__item"},{children:e.abbreviated+" "+e.average+" "}),i)}))}),void 0)]},void 0)}},4465:(e,i,a)=>{"use strict";a.r(i),a.d(i,{default:()=>m});var s=a(655),t=a(5893),r=a(9226),n=a(6317),l=a(3356),c=a(3727),o=function(e){var i=e.card;return(0,t.jsx)("div",(0,s.pi)({className:"marketplace-playerCard"},{children:(0,t.jsxs)(c.rU,(0,s.pi)({className:"marketplace-playerCard__link",to:{pathname:l.FM.FootballerCard.path,state:{card:i}}},{children:[(0,t.jsx)(n.s,{card:i,parentClassName:"marketplace-playerCard"},void 0),(0,t.jsxs)("div",(0,s.pi)({className:"marketplace-playerCard__price"},{children:[(0,t.jsx)("img",{className:"marketplace-playerCard__price__picture",src:i.mainInfo.priceIcon,alt:"Player price"},void 0),(0,t.jsx)("span",(0,s.pi)({className:"marketplace-playerCard__price__current"},{children:i.mainInfo.price}),void 0),(0,t.jsx)("img",{className:"marketplace-playerCard__price__status",src:i.mainInfo.priceStatus,alt:"Price status"},void 0)]}),void 0)]}),void 0)}),void 0)},d=function(e){var i=e.cards;return(0,t.jsx)("div",(0,s.pi)({className:"marketplace-cards"},{children:(0,t.jsx)("div",(0,s.pi)({className:"marketplace-cards__wrapper"},{children:i.map((function(e,i){return(0,t.jsx)(o,{card:e},i)}))}),void 0)}),void 0)},p=a(3382),u=a(7928);const m=function(){var e=(0,r.v9)((function(e){return e.cardsReducer.cards}));return(0,t.jsxs)("section",(0,s.pi)({className:"marketplace"},{children:[(0,t.jsx)(p.R,{title:"MARKETPLACE"},void 0),(0,t.jsx)(d,{cards:e},void 0),(0,t.jsx)(u.D,{itemCount:e.length},void 0)]}),void 0)}}}]);
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2FwcC9jb21wb25lbnRzL2NvbW1vbi9GaWx0ZXJGaWVsZC9GaWx0ZXJGaWVsZERyb3Bkb3duL2luZGV4LnRzeCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2FwcC9jb21wb25lbnRzL2NvbW1vbi9GaWx0ZXJGaWVsZC9pbmRleC50c3giLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC8uL3NyYy9hcHAvY29tcG9uZW50cy9jb21tb24vUGFnaW5hdG9yL1BhZ2luYXRvckJsb2NrUGFnZXMvaW5kZXgudHN4Iiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvLi9zcmMvYXBwL2NvbXBvbmVudHMvY29tbW9uL1BhZ2luYXRvci9pbmRleC50c3giLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC8uL3NyYy9hcHAvY29tcG9uZW50cy9jb21tb24vUGxheWVyQ2FyZC9pbmRleC50c3giLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC8uL3NyYy9hcHAvY29tcG9uZW50cy9NYXJrZXRQbGFjZS9NYXJrZXRQbGFjZUNhcmRzR3JvdXAvTWFya2V0UGxhY2VGb290YmFsbGVyQ2FyZC9pbmRleC50c3giLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC8uL3NyYy9hcHAvY29tcG9uZW50cy9NYXJrZXRQbGFjZS9NYXJrZXRQbGFjZUNhcmRzR3JvdXAvaW5kZXgudHN4Iiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvLi9zcmMvYXBwL3ZpZXdzL01hcmtldFBsYWNlUGFnZS9pbmRleC50c3giXSwibmFtZXMiOlsiRmlsdGVyRmllbGREcm9wZG93biIsInByb3BzIiwibGFiZWwiLCJpbWFnZSIsInVzZVN0YXRlIiwic2hvdWxkRHJvcGRvd25TaG93IiwiaGFuZGxlU2hvd2luZyIsImNsYXNzTmFtZSIsIm9uQ2xpY2siLCJwcmV2Iiwic3JjIiwiYWx0Iiwic3R5bGUiLCJkaXNwbGF5IiwiRmlsdGVyRmllbGQiLCJ0aXRsZSIsInNlYXJjaERhdGEiLCJzZXRTZWFyY2hEYXRhIiwiZmlsdGVyRmllbGRUaXRsZXMiLCJGaWx0ZXJJdGVtIiwicmVjdGFuZ2xlIiwic3RhciIsImZ1dCIsImV5ZSIsInBhcmFtZXRyZXMiLCJzZWFyY2giLCJ2YWx1ZSIsInBsYWNlaG9sZGVyIiwib25DaGFuZ2UiLCJldmVudCIsInRhcmdldCIsIm1hcCIsIml0ZW0iLCJpbmRleCIsIlBhZ2luYXRvckJsb2NrUGFnZXMiLCJibG9ja1BhZ2VzIiwib25QYWdlQ2hhbmdlIiwicGFnZSIsIlBhZ2luYXRvciIsIml0ZW1Db3VudCIsImN1cnJlbnRQYWdlIiwic2V0Q3VycmVudFBhZ2UiLCJmaXJzdEJsb2NrUGFnZXMiLCJzZXRGaXJzdEJsb2NrUGFnZXMiLCJtaWRkbGVCbG9ja1BhZ2VzIiwic2V0TWlkZGxlQmxvY2tQYWdlcyIsImxhc3RCbG9ja1BhZ2VzIiwic2V0TGFzdEJsb2NrUGFnZXMiLCJ1c2VFZmZlY3QiLCJwb3B1bGF0ZVBhZ2VzIiwicGFnZXMiLCJpIiwiTWF0aCIsImNlaWwiLCJwdXNoIiwiaXNGaXJzdERvdHNTaG93biIsImxlbmd0aCIsImlzU2Vjb25kRG90c1Nob3duIiwiaXNDdXJyZW50SW5GaXJzdEJsb2NrIiwiaXNDdXJyZW50SW5MYXN0QmxvY2siLCJNQVhfUEFHRVNfUEVSX0JMT0NLIiwiaXNPbmVCbG9ja1JlcXVpcmVkIiwic2xpY2UiLCJyZW9yZ2FuaXplUGFnZXNCbG9jayIsInR5cGUiLCJwYWdlTnVtYmVyIiwicHJldmlvdXMiLCJuZXh0IiwiUGxheWVyQ2FyZCIsImNhcmQiLCJwYXJlbnRDbGFzc05hbWUiLCJtYWluSW5mbyIsImJhY2tncm91bmRUeXBlIiwiZHJhZ2dhYmxlIiwicGxheWVyRmFjZSIsInRvIiwicGF0aG5hbWUiLCJzdGF0ZSIsImxhc3ROYW1lIiwic3RhdHMiLCJwcm9wZXJ0eSIsImFiYnJldmlhdGVkIiwiYXZlcmFnZSIsIk1hcmtldFBsYWNlRm9vdGJhbGxlckNhcmQiLCJwcmljZUljb24iLCJwcmljZSIsInByaWNlU3RhdHVzIiwiTWFya2V0UGxhY2VDYXJkc0dyb3VwIiwiY2FyZHMiLCJjYXJkc1JlZHVjZXIiXSwibWFwcGluZ3MiOiJxdkJBR2FBLEVBQTZFLFNBQUMsRyxJQUFFQyxFQUFLLFFBQ3RGQyxFQUFpQkQsRUFBSyxNQUFmRSxFQUFVRixFQUFLLE1BQ3hCLEdBQXNDLElBQUFHLFdBQVMsR0FBOUNDLEVBQWtCLEtBQUVDLEVBQWEsS0FFeEMsT0FDSSwwQkFDSUMsVUFBVSxjQUNWQyxRQUFTLFdBQU0sT0FBQUYsR0FBYyxTQUFBRyxHQUFRLE9BQUNBLE9BQUssWUFFM0MsMEJBQU1GLFVBQVUsc0JBQW9CLFVBQy9CTCxTQUFLLElBRVYsZ0JBQ0lLLFVBQVUsdUJBQ1ZHLElBQUtQLEVBQ0xRLElBQUtSLEdBQVMsb0JBQWEsSUFFL0IsZ0JBQ0lJLFVBQVUsd0JBQ1ZLLE1BQU8sQ0FBRUMsUUFBU1IsRUFBcUIsUUFBVSxjQUFRLFdBR3ZELElDSkxTLEVBQTJDLFNBQUMsRyxJQUFFQyxFQUFLLFFBQ3RELEdBQThCLElBQUFYLFVBQVMsSUFBdENZLEVBQVUsS0FBRUMsRUFBYSxLQUdoQyxFQUVJLFNBQ1dmLEVBQ0FDLEdBREEsS0FBQUQsUUFDQSxLQUFBQyxTQVFUZSxFQUFrQyxDQUNwQyxJQUFJQyxFQUFXLFVBQVdDLEdBQzFCLElBQUlELEVBQVcsWUFBYUMsR0FDNUIsSUFBSUQsRUFBVyxVQUFXQyxHQUMxQixJQUFJRCxFQUFXLFVBQVdDLEdBQzFCLElBQUlELEVBQVcsTUFBT0MsR0FDdEIsSUFBSUQsRUFBVyxRQUFTQyxHQUN4QixJQUFJRCxFQUFXLEdBQUlFLEdBQ25CLElBQUlGLEVBQVcsS0FBTUcsR0FDckIsSUFBSUgsRUFBVyxNQUFPQyxHQUN0QixJQUFJRCxFQUFXLEdBQUlJLEdBQ25CLElBQUlKLEVBQVcsR0FBSUUsR0FDbkIsSUFBSUYsRUFBVyxNQUFPQyxHQUN0QixJQUFJRCxFQUFXLEdBQUlLLEdBQ25CLElBQUlMLEVBQVcsT0FBUUMsSUFHM0IsT0FDSSw4QkFBU2IsVUFBVSxnQkFBYyxZ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c2VTdGF0ZSB9IGZyb20gJ3JlYWN0JztcbmltcG9ydCAnLi9pbmRleC5zY3NzJztcblxuZXhwb3J0IGNvbnN0IEZpbHRlckZpZWxkRHJvcGRvd246IFJlYWN0LkZDPHsgcHJvcHM6IHsgbGFiZWw6IHN0cmluZzsgaW1hZ2U6IHN0cmluZyB9IH0+ID0gKHsgcHJvcHMgfSkgPT4ge1xuICAgIGNvbnN0IHsgbGFiZWwsIGltYWdlIH0gPSBwcm9wcztcbiAgICBjb25zdCBbc2hvdWxkRHJvcGRvd25TaG93LCBoYW5kbGVTaG93aW5nXSA9IHVzZVN0YXRlKGZhbHNlKTtcblxuICAgIHJldHVybiAoXG4gICAgICAgIDxkaXZcbiAgICAgICAgICAgIGNsYXNzTmFtZT1cImZpbHRlci1pdGVtXCJcbiAgICAgICAgICAgIG9uQ2xpY2s9eygpID0+IGhhbmRsZVNob3dpbmcocHJldiA9PiAhcHJldil9XG4gICAgICAgID5cbiAgICAgICAgICAgIDxzcGFuIGNsYXNzTmFtZT1cImZpbHRlci1pdGVtX190aXRsZVwiPlxuICAgICAgICAgICAgICAgIHtsYWJlbH1cbiAgICAgICAgICAgIDwvc3Bhbj5cbiAgICAgICAgICAgIDxpbWdcbiAgICAgICAgICAgICAgICBjbGFzc05hbWU9XCJmaWx0ZXItaXRlbV9fcGljdHVyZVwiXG4gICAgICAgICAgICAgICAgc3JjPXtpbWFnZX1cbiAgICAgICAgICAgICAgICBhbHQ9e2ltYWdlICYmICdmaWx0ZXIgaWNvbid9XG4gICAgICAgICAgICAvPlxuICAgICAgICAgICAgPGRpdlxuICAgICAgICAgICAgICAgIGNsYXNzTmFtZT1cImZpbHRlci1pdGVtX19kcm9wZG93blwiXG4gICAgICAgICAgICAgICAgc3R5bGU9e3sgZGlzcGxheTogc2hvdWxkRHJvcGRvd25TaG93ID8gJ2Jsb2NrJyA6ICdub25lJyB9fVxuICAgICAgICAgICAgPlxuXG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgPC9kaXY+XG4gICAgKTtcbn07XG4iLCIvLyBDb3B5cmlnaHQgKEMpIDIwMjEgQ3JlZGl0b3IgQ29ycC4gR3JvdXAuXG4vLyBTZWUgTElDRU5TRSBmb3IgY29weWluZyBpbmZvcm1hdGlvbi5cblxuaW1wb3J0IHsgdXNlU3RhdGUgfSBmcm9tICdyZWFjdCc7XG5cbmltcG9ydCByZWN0YW5nbGVcbiAgICBmcm9tICdAc3RhdGljL2ltZy9GaWx0ZXJGaWVsZC9yZWN0YW5nbGUuc3ZnJztcbmltcG9ydCBzZWFyY2hcbiAgICBmcm9tICdAc3RhdGljL2ltZy9GaWx0ZXJGaWVsZC9zZWFyY2guc3ZnJztcbmltcG9ydCBzdGFyXG4gICAgZnJvbSAnQHN0YXRpYy9pbWcvRmlsdGVyRmllbGQvc3Rhci5zdmcnO1xuaW1wb3J0IGZ1dFxuICAgIGZyb20gJ0BzdGF0aWMvaW1nL0ZpbHRlckZpZWxkL2Z1dC5zdmcnO1xuaW1wb3J0IGV5ZVxuICAgIGZyb20gJ0BzdGF0aWMvaW1nL0ZpbHRlckZpZWxkL2V5ZS5zdmcnO1xuaW1wb3J0IHBhcmFtZXRyZXNcbiAgICBmcm9tICdAc3RhdGljL2ltZy9GaWx0ZXJGaWVsZC9wYXJhbWV0cmVzLnN2Zyc7XG5cbmltcG9ydCAnLi9pbmRleC5zY3NzJztcbmltcG9ydCB7IEZpbHRlckZpZWxkRHJvcGRvd24gfSBmcm9tICcuL0ZpbHRlckZpZWxkRHJvcGRvd24nO1xuXG5leHBvcnQgY29uc3QgRmlsdGVyRmllbGQ6IFJlYWN0LkZDPHsgdGl0bGU6IHN0cmluZyB9PiA9ICh7IHRpdGxlIH0pID0+IHtcbiAgICBjb25zdCBbc2VhcmNoRGF0YSwgc2V0U2VhcmNoRGF0YV0gPSB1c2VTdGF0ZSgnJyk7XG5cbiAgICAvKiogQ2xhc3MgZm9yIGVhY2ggZmlsdGVyIGZpZWxkIGl0ZW0gKi9cbiAgICBjbGFzcyBGaWx0ZXJJdGVtIHtcbiAgICAgICAgLyoqIGxhYmUgaXMgdGV4dCBmaWxsaW5nIG9mIGZpZWxkLCBpbWFnZSAtIHNyYyBmb3IgaW1hZ2UgICovXG4gICAgICAgIGNvbnN0cnVjdG9yKFxuICAgICAgICAgICAgcHVibGljIGxhYmVsOiBzdHJpbmcsXG4gICAgICAgICAgICBwdWJsaWMgaW1hZ2U6IHN0cmluZyxcbiAgICAgICAgKSB7IH1cbiAgICB9XG5cbiAgICBjb25zdCBoYW5kbGVTZXJjaENoYW5nZSA9IChldmVudDogYW55KSA9PiB7XG4gICAgICAgIHNldFNlYXJjaERhdGEoZXZlbnQudGFyZ2V0LnZhbHVlKTtcbiAgICB9O1xuXG4gICAgY29uc3QgZmlsdGVyRmllbGRUaXRsZXM6IEZpbHRlckl0ZW1bXSA9IFtcbiAgICAgICAgbmV3IEZpbHRlckl0ZW0oJ1ZlcnNpb24nLCByZWN0YW5nbGUpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnUG9zaXRpb25zJywgcmVjdGFuZ2xlKSxcbiAgICAgICAgbmV3IEZpbHRlckl0ZW0oJ05hdGlvbnMnLCByZWN0YW5nbGUpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnTGVhZ3VlcycsIHJlY3RhbmdsZSksXG4gICAgICAgIG5ldyBGaWx0ZXJJdGVtKCdXUkYnLCByZWN0YW5nbGUpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnU3RhdHMnLCByZWN0YW5nbGUpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnJywgc3RhciksXG4gICAgICAgIG5ldyBGaWx0ZXJJdGVtKCdQUycsIGZ1dCksXG4gICAgICAgIG5ldyBGaWx0ZXJJdGVtKCdUJlMnLCByZWN0YW5nbGUpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnJywgZXllKSxcbiAgICAgICAgbmV3IEZpbHRlckl0ZW0oJycsIHN0YXIpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnUlBQJywgcmVjdGFuZ2xlKSxcbiAgICAgICAgbmV3IEZpbHRlckl0ZW0oJycsIHBhcmFtZXRyZXMpLFxuICAgICAgICBuZXcgRmlsdGVySXRlbSgnTWlzYycsIHJlY3RhbmdsZSksXG4gICAgXTtcblxuICAgIHJldHVybiAoXG4gICAgICAgIDxzZWN0aW9uIGNsYXNzTmFtZT1cImZpbHRlci1maWVsZFwiPlxuICAgICAgICAgICAgPGgxIGNsYXNzTmFtZT1cImZpbHRlci1maWVsZF9fdGl0bGVcIj5cbiAgICAgICAgICAgICAgICB7dGl0bGV9XG4gICAgICAgICAgICA8L2gxPlxuICAgICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJmaWx0ZXItZmllbGRfX3dyYXBwZXJcIj5cbiAgICAgICAgICAgICAgICA8dWwgY2xhc3NOYW1lPVwiZmlsdGVyLWZpZWxkX19saXN0XCI+XG4gICAgICAgICAgICAgICAgICAgIDxsaSBjbGFzc05hbWU9XCJmaWx0ZXItZmllbGRfX2xpc3RfX2l0ZW1cIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxpbWdcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBzcmM9e3NlYXJjaH1cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBhbHQ9XCJGaWx0ZXIgSWNvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgY2xhc3NOYW1lPVwiZmlsdGVyLWZpZWxkX19saXN0X19pdGVtX19waWN0dXJlXCJcbiAgICAgICAgICAgICAgICAgICAgICAgIC8+XG4gICAgICAgICAgICAgICAgICAgICAgICA8aW5wdXRcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YWx1ZT17c2VhcmNoRGF0YX1cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBwbGFjZWhvbGRlcj1cIlNlYXJjaFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgY2xhc3NOYW1lPVwiZmlsdGVyLWZpZWxkX19saXN0X19pdGVtX19zZWFyY2hcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIG9uQ2hhbmdlPXtoYW5kbGVTZXJjaENoYW5nZX1cbiAgICAgICAgICAgICAgICAgICAgICAgIC8+XG4gICAgICAgICAgICAgICAgICAgIDwvbGk+XG4gICAgICAgICAgICAgICAgICAgIHtmaWx0ZXJGaWVsZFRpdGxlcy5tYXAoKGl0ZW0sIGluZGV4KSA9PlxuICAgICAgICAgICAgICAgICAgICAgICAgPGxpXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgY2xhc3NOYW1lPVwiZmlsdGVyLWZpZWxkX19saXN0X19pdGVtXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBrZXk9e2luZGV4fVxuICAgICAgICAgICAgICAgICAgICAgICAgPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxGaWx0ZXJGaWVsZERyb3Bkb3duIHByb3BzPXtpdGVtfSAvPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9saT5cbiAgICAgICAgICAgICAgICAgICAgKX1cbiAgICAgICAgICAgICAgICA8L3VsPlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgIDwvc2VjdGlvbiA+XG4gICAgKTtcbn07XG4iLCIvLyBDb3B5cmlnaHQgKEMpIDIwMjEgQ3JlZGl0b3IgQ29ycC4gR3JvdXAuXG4vLyBTZWUgTElDRU5TRSBmb3IgY29weWluZyBpbmZvcm1hdGlvbi5cblxuLyogZXNsaW50LWRpc2FibGUgKi9cbmV4cG9ydCBjb25zdCBQYWdpbmF0b3JCbG9ja1BhZ2VzOiBSZWFjdC5GQzx7XG4gICAgYmxvY2tQYWdlczogbnVtYmVyW107XG4gICAgb25QYWdlQ2hhbmdlOiAodHlwZTogc3RyaW5nLCBwYWdlTnVtYmVyPzogbnVtYmVyKSA9PiB2b2lkO1xufT4gPSAoeyBibG9ja1BhZ2VzLCBvblBhZ2VDaGFuZ2UgfSkgPT5cbiAgICAgICAgPHVsIGNsYXNzTmFtZT1cInVsdGltYXRlZGl2aXNpb24tcGFnaW5hdG9yX19wYWdlc1wiPlxuICAgICAgICAgICAge2Jsb2NrUGFnZXMubWFwKChwYWdlLCBpbmRleCkgPT5cbiAgICAgICAgICAgICAgICA8bGlcbiAgICAgICAgICAgICAgICAgICAgY2xhc3NOYW1lPVwidWx0aW1hdGVkaXZpc2lvbi1wYWdpbmF0b3JfX3BhZ2VzX19pdGVtXCJcbiAgICAgICAgICAgICAgICAgICAga2V5PXtpbmRleH1cbiAgICAgICAgICAgICAgICAgICAgb25DbGljaz17KCkgPT4gb25QYWdlQ2hhbmdlKCdjaGFuZ2UgcGFnZScsIHBhZ2UpfVxuICAgICAgICAgICAgICAgID5cbiAgICAgICAgICAgICAgICAgICAge3BhZ2V9XG4gICAgICAgICAgICAgICAgPC9saT4sXG4gICAgICAgICAgICApfVxuICAgICAgICA8L3VsPlxuICAgIDtcbiIsIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuLyogZXNsaW50LWRpc2FibGUgKi9cbmltcG9ydCB7IHVzZUVmZmVjdCwgdXNlU3RhdGUgfSBmcm9tICdyZWFjdCc7XG5cbmltcG9ydCBuZXh0IGZyb20gJ0BzdGF0aWMvaW1nL1VsdGltYXRlRGl2aXNpb25QYWdpbmF0b3IvbmV4dC5zdmcnO1xuaW1wb3J0IHByZXZpb3VzIGZyb20gJ0BzdGF0aWMvaW1nL1VsdGltYXRlRGl2aXNpb25QYWdpbmF0b3IvcHJldmlvdXMuc3ZnJztcbmltcG9ydCB7IFBhZ2luYXRvckJsb2NrUGFnZXMgfSBmcm9tICdAY29tcG9uZW50cy9jb21tb24vUGFnaW5hdG9yL1BhZ2luYXRvckJsb2NrUGFnZXMnO1xuXG5pbXBvcnQgJy4vaW5kZXguc2Nzcyc7XG5cbmV4cG9ydCBjb25zdCBQYWdpbmF0b3I6IFJlYWN0LkZDPHsgaXRlbUNvdW50OiBudW1iZXIgfT4gPSAoeyBpdGVtQ291bnQgfSkgPT4ge1xuICAgIGNvbnN0IEZJUlNUX0lURU1fUEFHSU5BVE9OID0gMTtcbiAgICBjb25zdCBbY3VycmVudFBhZ2UsIHNldEN1cnJlbnRQYWdlXSA9IHVzZVN0YXRlPG51bWJlcj4oRklSU1RfSVRFTV9QQUdJTkFUT04pO1xuICAgIC8qKlxuICAgICogc3BsaXQgdGhlIHBhZ2UgaW50byAzIGJsb2NrcyB0aGF0IGNhbiBiZSBuZWVkZWRcbiAgICAqIHRvIHNlcGFyYXRlIHBhZ2UgbnVtYmVyc1xuICAgICAqL1xuICAgIGNvbnN0IFtmaXJzdEJsb2NrUGFnZXMsIHNldEZpcnN0QmxvY2tQYWdlc10gPSB1c2VTdGF0ZTxudW1iZXJbXT4oW10pO1xuICAgIGNvbnN0IFttaWRkbGVCbG9ja1BhZ2VzLCBzZXRNaWRkbGVCbG9ja1BhZ2VzXSA9IHVzZVN0YXRlPG51bWJlcltdPihbXSk7XG4gICAgY29uc3QgW2xhc3RCbG9ja1BhZ2VzLCBzZXRMYXN0QmxvY2tQYWdlc10gPSB1c2VTdGF0ZTxudW1iZXJbXT4oW10pO1xuXG4gICAgdXNlRWZmZWN0KCgpID0+IHtcbiAgICAgICAgcG9wdWxhdGVQYWdlcygpO1xuICAgIH0sIFtjdXJyZW50UGFnZV0pO1xuXG4gICAgY29uc3QgQ0FSRFNfT05fUEFHRTogbnVtYmVyID0gODtcbiAgICBjb25zdCBNQVhfUEFHRVNfUEVSX0JMT0NLOiBudW1iZXIgPSA1O1xuICAgIGNvbnN0IE1BWF9QQUdFU19PRkZfQkxPQ0tTOiBudW1iZXIgPSAxMDtcbiAgICBjb25zdCBGSVJTVF9QQUdFX0lOREVYOiBudW1iZXIgPSAwO1xuICAgIGNvbnN0IFNFQ09ORF9QQUdFX0lOREVYOiBudW1iZXIgPSAxO1xuICAgIGNvbnN0IEZJUlNUX1BBR0VfSU5ERVhfRlJPTV9FTkQ6IG51bWJlciA9IC0xO1xuICAgIGNvbnN0IE5FR19TVEVQX0ZST01fQ1VSUkVOVF9QQUdFOiBudW1iZXIgPSAtMztcbiAgICBjb25zdCBQT1NfU1RFUF9GUk9NX0NVUlJFTlRfUEFHRTogbnVtYmVyID0gMjtcblxuICAgIGNvbnN0IHBhZ2VzOiBudW1iZXJbXSA9IFtdO1xuICAgIGZvciAobGV0IGkgPSAxOyBpIDw9IE1hdGguY2VpbChpdGVtQ291bnQgLyBDQVJEU19PTl9QQUdFKTsgaSsrKSB7XG4gICAgICAgIHBhZ2VzLnB1c2goaSk7XG4gICAgfVxuICAgIC8qKiBzZXQgYmxvY2sgcGFnZXMgZGVwZW5kcyBvbiBjdXJyZW50IHBhZ2UgKi9cbiAgICBjb25zdCBzZXRCbG9ja3NJZkN1cnJlbnRJbkZpcnN0QmxvY2sgPSAoKSA9PiB7XG4gICAgICAgIHNldEZpcnN0QmxvY2tQYWdlcyhwYWdlcy5zbGljZShGSVJTVF9QQUdFX0lOREVYLCBNQVhfUEFHRVNfUEVSX0JMT0NLKSk7XG4gICAgICAgIHNldE1pZGRsZUJsb2NrUGFnZXMoW10pO1xuICAgICAgICBzZXRMYXN0QmxvY2tQYWdlcyhwYWdlcy5zbGljZShGSVJTVF9QQUdFX0lOREVYX0ZST01fRU5EKSk7XG4gICAgfTtcbiAgICBjb25zdCBzZXRCbG9ja3NJZkN1cnJlbnRJbk1pZGRsZUJsb2NrID0gKCkgPT4ge1xuICAgICAgICBzZXRGaXJzdEJsb2NrUGFnZXMocGFnZXMuc2xpY2UoRklSU1RfUEFHRV9JTkRFWCwgU0VDT05EX1BBR0VfSU5ERVgpKTtcbiAgICAgICAgc2V0TWlkZGxlQmxvY2tQYWdlcyhwYWdlcy5zbGljZShjdXJyZW50UGFnZSArIE5FR19TVEVQX0ZST01fQ1VSUkVOVF9QQUdFLCBjdXJyZW50UGFnZSArIFBPU19TVEVQX0ZST01fQ1VSUkVOVF9QQUdFKSk7XG4gICAgICAgIHNldExhc3RCbG9ja1BhZ2VzKHBhZ2VzLnNsaWNlKEZJUlNUX1BBR0VfSU5ERVhfRlJPTV9FTkQpKTtcbiAgICB9O1xuICAgIGNvbnN0IHNldEJsb2Nrc0lmQ3VycmVudEluTGFzdEJsb2NrID0gKCkgPT4ge1xuICAgICAgICBzZXRGaXJzdEJsb2NrUGFnZXMocGFnZXMuc2xpY2UoRklSU1RfUEFHRV9JTkRFWCwgU0VDT05EX1BBR0VfSU5ERVgpKTtcbiAgICAgICAgc2V0TWlkZGxlQmxvY2tQYWdlcyhbXSk7XG4gICAgICAgIHNldExhc3RCbG9ja1BhZ2VzKHBhZ2VzLnNsaWNlKC1NQVhfUEFHRVNfUEVSX0JMT0NLKSk7XG4gICAgfTtcbiAgICAvKipcbiAgICAqIEluZGljYXRlcyB2aXNpYmlsaXR5IG9mIGRvdHMgYWZ0ZXIgZmlyc3QgcGFnZXMgYmxvY2tcbiAgICAgKi9cbiAgICBjb25zdCBpc0ZpcnN0RG90c1Nob3duOiBib29sZWFuID1cbiAgICAgICAgbWlkZGxlQmxvY2tQYWdlcy5sZW5ndGggPD0gTUFYX1BBR0VTX1BFUl9CTE9DS1xuICAgICAgICAmJiBwYWdlcy5sZW5ndGggPiBNQVhfUEFHRVNfT0ZGX0JMT0NLUztcbiAgICAvKlxuICAgICogSW5kaWNhdGVzIHZpc2liaWxpdHkgb2YgZG90cyBhZnRlciBtaWRkbGUgcGFnZXMgYmxvY2tcbiAgICAgKi9cbiAgICBjb25zdCBpc1NlY29uZERvdHNTaG93bjogYm9vbGVhbiA9ICEhbWlkZGxlQmxvY2tQYWdlcy5sZW5ndGg7XG4gICAgLyoqXG4gICAgICogaW5kaWNhdGVzIGluIHdoaWNoIGJsb2NrIGN1cnJlbnQgcGFnZVxuICAgICAqL1xuICAgIGNvbnN0IGlzQ3VycmVudEluRmlyc3RCbG9jazogYm9vbGVhbiA9IGN1cnJlbnRQYWdlIDwgTUFYX1BBR0VTX1BFUl9CTE9DSztcbiAgICBjb25zdCBpc0N1cnJlbnRJbkxhc3RCbG9jazogYm9vbGVhbiA9IHBhZ2VzLmxlbmd0aCAtIGN1cnJlbnRQYWdlIDwgTUFYX1BBR0VTX1BFUl9CTE9DSyAtIFNFQ09ORF9QQUdFX0lOREVYO1xuICAgIC8qKlxuICAgICAqIGNoYW5nZSBwYWdlIGJsb2NrcyByZW9yZ2FuaXphdGlvbiBkZXBlbmRzXG4gICAgICogb24gY3VycmVudCBwYWdlXG4gICAgICovXG4gICAgY29uc3QgcmVvcmdhbml6ZVBhZ2VzQmxvY2sgPSAoKSA9PiB7XG4gICAgICAgIGlmIChpc09uZUJsb2NrUmVxdWlyZWQpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBpZiAoaXNDdXJyZW50SW5GaXJzdEJsb2NrKSB7XG4gICAgICAgICAgICBzZXRCbG9ja3NJZkN1cnJlbnRJbkZpcnN0QmxvY2soKTtcblxuICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICB9XG4gICAgICAgIGlmICghaXNDdXJyZW50SW5GaXJzdEJsb2NrICYmICFpc0N1cnJlbnRJbkxhc3RCbG9jaykge1xuICAgICAgICAgICAgc2V0QmxvY2tzSWZDdXJyZW50SW5NaWRkbGVCbG9jaygpO1xuXG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgaWYgKGlzQ3VycmVudEluTGFzdEJsb2NrKSB7XG4gICAgICAgICAgICBzZXRCbG9ja3NJZkN1cnJlbnRJbkxhc3RCbG9jaygpO1xuICAgICAgICB9XG4gICAgfTtcbiAgICAvKlxuICAgICogaW5kaWNhdGVzIGlmIGRvdHMgZGVsaW1pdGVyIGlzIG5lZWRlZFxuICAgICogdG8gc2VwYXJhdGUgcGFnZSBudW1iZXJzXG4gICAgKi9cbiAgICBjb25zdCBpc09uZUJsb2NrUmVxdWlyZWQ6IGJvb2xlYW4gPSBwYWdlcy5sZW5ndGggPD0gTUFYX1BBR0VTX09GRl9CTE9DS1M7XG4gICAgY29uc3QgcG9wdWxhdGVQYWdlcyA9ICgpID0+IHtcbiAgICAgICAgaWYgKCFwYWdlcy5sZW5ndGgpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBpZiAoaXNPbmVCbG9ja1JlcXVpcmVkKSB7XG4gICAgICAgICAgICBzZXRGaXJzdEJsb2NrUGFnZXMocGFnZXMuc2xpY2UoKSk7XG4gICAgICAgICAgICBzZXRNaWRkbGVCbG9ja1BhZ2VzKFtdKTtcbiAgICAgICAgICAgIHNldExhc3RCbG9ja1BhZ2VzKFtdKTtcblxuICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICB9XG4gICAgICAgIHJlb3JnYW5pemVQYWdlc0Jsb2NrKCk7XG4gICAgfTtcbiAgICAvKipcbiAgICAgKiBjaGFuZ2UgY3VycmVudCBwYWdlIGFuZCBzZXQgcGFnZXMgYmxvY2tcbiAgICAgKi9cbiAgICBjb25zdCBvblBhZ2VDaGFuZ2UgPSAodHlwZTogc3RyaW5nLCBwYWdlTnVtYmVyOiBudW1iZXIgPSBjdXJyZW50UGFnZSk6IHZvaWQgPT4ge1xuICAgICAgICBjb25zdCBTVEVQX0ZST01fQ1VSUkVOVF9QQUdFID0gMTtcbiAgICAgICAgc3dpdGNoICh0eXBlKSB7XG4gICAgICAgICAgICBjYXNlICduZXh0IHBhZ2UnOlxuICAgICAgICAgICAgICAgIGlmIChwYWdlTnVtYmVyIDwgcGFnZXMubGVuZ3RoKSB7XG4gICAgICAgICAgICAgICAgICAgIHNldEN1cnJlbnRQYWdlKHBhZ2VOdW1iZXIgKyBTVEVQX0ZST01fQ1VSUkVOVF9QQUdFKTtcbiAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICAgICAgcG9wdWxhdGVQYWdlcygpO1xuXG4gICAgICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICAgICAgY2FzZSAncHJldmlvdXMgcGFnZSc6XG4gICAgICAgICAgICAgICAgaWYgKHBhZ2VOdW1iZXIgPiBTRUNPTkRfUEFHRV9JTkRFWCkge1xuICAgICAgICAgICAgICAgICAgICBzZXRDdXJyZW50UGFnZShwYWdlTnVtYmVyIC0gU1RFUF9GUk9NX0NVUlJFTlRfUEFHRSk7XG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgIHBvcHVsYXRlUGFnZXMoKTtcblxuICAgICAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgICAgIGNhc2UgJ2NoYW5nZSBwYWdlJzpcbiAgICAgICAgICAgICAgICBzZXRDdXJyZW50UGFnZShwYWdlTnVtYmVyKTtcbiAgICAgICAgICAgICAgICBwb3B1bGF0ZVBhZ2VzKCk7XG5cbiAgICAgICAgICAgICAgICByZXR1cm47XG4gICAgICAgICAgICBkZWZhdWx0OlxuICAgICAgICAgICAgICAgIHBvcHVsYXRlUGFnZXMoKTtcbiAgICAgICAgfVxuICAgIH07XG5cbiAgICByZXR1cm4gKFxuICAgICAgICA8c2VjdGlvbiBjbGFzc05hbWU9XCJ1bHRpbWF0ZWRpdmlzaW9uLXBhZ2luYXRvclwiPlxuICAgICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJ1bHRpbWF0ZWRpdmlzaW9uLXBhZ2luYXRvcl9fd3JhcHBlclwiPlxuICAgICAgICAgICAgICAgIDxhIGNsYXNzTmFtZT1cInVsdGltYXRlZGl2aXNpb24tcGFnaW5hdG9yX19wcmV2aW91c1wiXG4gICAgICAgICAgICAgICAgICAgIG9uQ2xpY2s9eygpID0+IG9uUGFnZUNoYW5nZSgncHJldmlvdXMgcGFnZScpfT5cbiAgICAgICAgICAgICAgICAgICAgPGltZyBjbGFzc05hbWU9XCJ1bHRpbWF0ZWRpdmlzaW9uLXBhZ2luYXRvcl9fcHJldmlvdXNfX2Fycm93XCJcbiAgICAgICAgICAgICAgICAgICAgICAgIHNyYz17cHJldmlvdXN9XG4gICAgICAgICAgICAgICAgICAgICAgICBhbHQ9XCJQcmV2aW91cyBwYWdlXCIgLz5cbiAgICAgICAgICAgICAgICAgICAgPHAgY2xhc3NOYW1lPVwidWx0aW1hdGVkaXZpc2lvbi1wYWdpbmF0b3JfX3ByZXZpb3VzX190aXRsZVwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgUHJldmlvdXMgcGFnZVxuICAgICAgICAgICAgICAgICAgICA8L3A+XG4gICAgICAgICAgICAgICAgPC9hPlxuICAgICAgICAgICAgICAgIDxQYWdpbmF0b3JCbG9ja1BhZ2VzXG4gICAgICAgICAgICAgICAgICAgIGJsb2NrUGFnZXM9e2ZpcnN0QmxvY2tQYWdlc31cbiAgICAgICAgICAgICAgICAgICAgb25QYWdlQ2hhbmdlPXtvblBhZ2VDaGFuZ2V9XG4gICAgICAgICAgICAgICAgLz5cbiAgICAgICAgICAgICAgICB7aXNGaXJzdERvdHNTaG93blxuICAgICAgICAgICAgICAgICAgICAmJiA8c3BhbiBjbGFzc05hbWU9XCJ1bHRpbWF0ZWRpdmlzaW9uLXBhZ2luYXRvcl9fcGFnZXNfX2RvdHNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIC4uLjwvc3Bhbj59XG4gICAgICAgICAgICAgICAgPFBhZ2luYXRvckJsb2NrUGFnZXNcbiAgICAgICAgICAgICAgICAgICAgYmxvY2tQYWdlcz17bWlkZGxlQmxvY2tQYWdlc31cbiAgICAgICAgICAgICAgICAgICAgb25QYWdlQ2hhbmdlPXtvblBhZ2VDaGFuZ2V9XG4gICAgICAgICAgICAgICAgLz5cbiAgICAgICAgICAgICAgICB7aXNTZWNvbmREb3RzU2hvd25cbiAgICAgICAgICAgICAgICAgICAgJiYgPHNwYW4gY2xhc3NOYW1lPVwidWx0aW1hdGVkaXZpc2lvbi1wYWdpbmF0b3JfX3BhZ2VzX19kb3RzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAuLi48L3NwYW4+fVxuICAgICAgICAgICAgICAgIDxQYWdpbmF0b3JCbG9ja1BhZ2VzXG4gICAgICAgICAgICAgICAgICAgIGJsb2NrUGFnZXM9e2xhc3RCbG9ja1BhZ2VzfVxuICAgICAgICAgICAgICAgICAgICBvblBhZ2VDaGFuZ2U9e29uUGFnZUNoYW5nZX1cbiAgICAgICAgICAgICAgICAvPlxuICAgICAgICAgICAgICAgIDxhIGNsYXNzTmFtZT1cInVsdGltYXRlZGl2aXNpb24tcGFnaW5hdG9yX19uZXh0XCJcbiAgICAgICAgICAgICAgICAgICAgb25DbGljaz17KCkgPT4gb25QYWdlQ2hhbmdlKCduZXh0IHBhZ2UnKX0+XG4gICAgICAgICAgICAgICAgICAgIDxwIGNsYXNzTmFtZT1cInVsdGltYXRlZGl2aXNpb24tcGFnaW5hdG9yX19uZXh0X190aXRsZVwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgTmV4dCBwYWdlXG4gICAgICAgICAgICAgICAgICAgIDwvcD5cbiAgICAgICAgICAgICAgICAgICAgPGltZyBjbGFzc05hbWU9XCJ1bHRpbWF0ZWRpdmlzaW9uLXBhZ2luYXRvcl9fbmV4dF9fYXJyb3dcIlxuICAgICAgICAgICAgICAgICAgICAgICAgc3JjPXtuZXh0fVxuICAgICAgICAgICAgICAgICAgICAgICAgYWx0PVwiTmV4dCBwYWdlXCIgLz5cbiAgICAgICAgICAgICAgICA8L2E+XG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgPC9zZWN0aW9uPlxuICAgICk7XG59O1xuIiwiLy8gQ29weXJpZ2h0IChDKSAyMDIxIENyZWRpdG9yIENvcnAuIEdyb3VwLlxuLy8gU2VlIExJQ0VOU0UgZm9yIGNvcHlpbmcgaW5mb3JtYXRpb24uXG5cbmltcG9ydCB7IExpbmssIE5hdkxpbmsgfSBmcm9tICdyZWFjdC1yb3V0ZXItZG9tJztcblxuaW1wb3J0IHsgUm91dGVDb25maWcgfSBmcm9tICdAL2FwcC9yb3V0ZXInO1xuXG5pbXBvcnQgeyBDYXJkIH0gZnJvbSAnQC9jYXJkJztcblxuZXhwb3J0IGNvbnN0IFBsYXllckNhcmQ6IFJlYWN0LkZDPHsgY2FyZDogQ2FyZDsgcGFyZW50Q2xhc3NOYW1lOiBzdHJpbmcgfT4gPSAoe1xuICAgIGNhcmQsIHBhcmVudENsYXNzTmFtZSxcbn0pID0+XG4gICAgPD5cbiAgICAgICAgPGltZ1xuICAgICAgICAgICAgY2xhc3NOYW1lPXtgJHtwYXJlbnRDbGFzc05hbWV9X19iYWNrZ3JvdW5kLXR5cGVgfVxuICAgICAgICAgICAgc3JjPXtjYXJkLm1haW5JbmZvLmJhY2tncm91bmRUeXBlfVxuICAgICAgICAgICAgYWx0PVwiYmFja2dyb3VuZCBpbWdcIlxuICAgICAgICAgICAgZHJhZ2dhYmxlPXtmYWxzZX1cbiAgICAgICAgLz5cbiAgICAgICAgPGltZ1xuICAgICAgICAgICAgY2xhc3NOYW1lPXtgJHtwYXJlbnRDbGFzc05hbWV9X19mYWNlLXBpY3R1cmVgfVxuICAgICAgICAgICAgc3JjPXtjYXJkLm1haW5JbmZvLnBsYXllckZhY2V9XG4gICAgICAgICAgICBhbHQ9XCJQbGF5ZXIgZmFjZVwiXG4gICAgICAgICAgICBkcmFnZ2FibGU9e2ZhbHNlfVxuICAgICAgICAvPlxuICAgICAgICA8TGlua1xuICAgICAgICAgICAgdG89e3tcbiAgICAgICAgICAgICAgICBwYXRobmFtZTogUm91dGVDb25maWcuRm9vdGJhbGxlckNhcmQucGF0aCxcbiAgICAgICAgICAgICAgICBzdGF0ZToge1xuICAgICAgICAgICAgICAgICAgICBjYXJkLFxuICAgICAgICAgICAgICAgIH0sXG4gICAgICAgICAgICB9fVxuICAgICAgICA+XG4gICAgICAgICAgICA8c3BhbiBjbGFzc05hbWU9e2Ake3BhcmVudENsYXNzTmFtZX1fX25hbWVgfT5cbiAgICAgICAgICAgICAgICB7Y2FyZC5tYWluSW5mby5sYXN0TmFtZX1cbiAgICAgICAgICAgIDwvc3Bhbj5cbiAgICAgICAgPC9MaW5rPlxuICAgICAgICA8dWwgY2xhc3NOYW1lPXtgJHtwYXJlbnRDbGFzc05hbWV9X19saXN0YH0+XG4gICAgICAgICAgICB7Y2FyZC5zdGF0cy5tYXAoXG4gICAgICAgICAgICAgICAgKHByb3BlcnR5LCBpbmRleCkgPT5cbiAgICAgICAgICAgICAgICAgICAgPGxpXG4gICAgICAgICAgICAgICAgICAgICAgICBjbGFzc05hbWU9e2Ake3BhcmVudENsYXNzTmFtZX1fX2xpc3RfX2l0ZW1gfVxuICAgICAgICAgICAgICAgICAgICAgICAga2V5PXtpbmRleH0+XG4gICAgICAgICAgICAgICAgICAgICAgICB7XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgLyoqXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICogZ2V0IG9ubHkgYXZlcmFnZSB2YWx1ZSBvZiBwbGF5ZXIncyBnYW1lIHByb3BlcnR5XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICovXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgYCR7cHJvcGVydHkuYWJicmV2aWF0ZWR9ICR7cHJvcGVydHkuYXZlcmFnZX0gYFxuICAgICAgICAgICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgICAgICA8L2xpPixcbiAgICAgICAgICAgICl9XG4gICAgICAgIDwvdWw+XG4gICAgPC8+O1xuIiwiLy8gQ29weXJpZ2h0IChDKSAyMDIxIENyZWRpdG9yIENvcnAuIEdyb3VwLlxuLy8gU2VlIExJQ0VOU0UgZm9yIGNvcHlpbmcgaW5mb3JtYXRpb24uXG5cbmltcG9ydCB7IFBsYXllckNhcmQgfSBmcm9tICdAY29tcG9uZW50cy9jb21tb24vUGxheWVyQ2FyZCc7XG5cbmltcG9ydCB7IFJvdXRlQ29uZmlnIH0gZnJvbSAnQC9hcHAvcm91dGVyJztcbmltcG9ydCB7IExpbmsgfSBmcm9tICdyZWFjdC1yb3V0ZXItZG9tJztcblxuaW1wb3J0IHsgQ2FyZCB9IGZyb20gJ0AvY2FyZCc7XG5cbmltcG9ydCAnLi9pbmRleC5zY3NzJztcblxuZXhwb3J0IGNvbnN0IE1hcmtldFBsYWNlRm9vdGJhbGxlckNhcmQ6IFJlYWN0LkZDPHsgY2FyZDogQ2FyZDsgcGxhY2U/OiBzdHJpbmcgfT4gPSAoeyBjYXJkIH0pID0+XG4gICAgPGRpdlxuICAgICAgICBjbGFzc05hbWU9XCJtYXJrZXRwbGFjZS1wbGF5ZXJDYXJkXCJcbiAgICA+XG4gICAgICAgIDxMaW5rXG4gICAgICAgICAgICBjbGFzc05hbWU9XCJtYXJrZXRwbGFjZS1wbGF5ZXJDYXJkX19saW5rXCJcbiAgICAgICAgICAgIHRvPXt7XG4gICAgICAgICAgICAgICAgcGF0aG5hbWU6IFJvdXRlQ29uZmlnLkZvb3RiYWxsZXJDYXJkLnBhdGgsXG4gICAgICAgICAgICAgICAgc3RhdGU6IHtcbiAgICAgICAgICAgICAgICAgICAgY2FyZCxcbiAgICAgICAgICAgICAgICB9LFxuICAgICAgICAgICAgfX1cbiAgICAgICAgPlxuICAgICAgICAgICAgPFBsYXllckNhcmRcbiAgICAgICAgICAgICAgICBjYXJkPXtjYXJkfVxuICAgICAgICAgICAgICAgIHBhcmVudENsYXNzTmFtZT17J21hcmtldHBsYWNlLXBsYXllckNhcmQnfVxuICAgICAgICAgICAgLz5cbiAgICAgICAgICAgIDxkaXYgY2xhc3NOYW1lPVwibWFya2V0cGxhY2UtcGxheWVyQ2FyZF9fcHJpY2VcIj5cbiAgICAgICAgICAgICAgICA8aW1nIGNsYXNzTmFtZT1cIm1hcmtldHBsYWNlLXBsYXllckNhcmRfX3ByaWNlX19waWN0dXJlXCJcbiAgICAgICAgICAgICAgICAgICAgc3JjPXtjYXJkLm1haW5JbmZvLnByaWNlSWNvbn1cbiAgICAgICAgICAgICAgICAgICAgYWx0PVwiUGxheWVyIHByaWNlXCIgLz5cbiAgICAgICAgICAgICAgICA8c3BhbiBjbGFzc05hbWU9XCJtYXJrZXRwbGFjZS1wbGF5ZXJDYXJkX19wcmljZV9fY3VycmVudFwiPlxuICAgICAgICAgICAgICAgICAgICB7Y2FyZC5tYWluSW5mby5wcmljZX1cbiAgICAgICAgICAgICAgICA8L3NwYW4+XG4gICAgICAgICAgICAgICAgPGltZyBjbGFzc05hbWU9XCJtYXJrZXRwbGFjZS1wbGF5ZXJDYXJkX19wcmljZV9fc3RhdHVzXCJcbiAgICAgICAgICAgICAgICAgICAgc3JjPXtjYXJkLm1haW5JbmZvLnByaWNlU3RhdHVzfVxuICAgICAgICAgICAgICAgICAgICBhbHQ9XCJQcmljZSBzdGF0dXNcIiAvPlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgIDwvTGluaz5cbiAgICA8L2RpdiA+O1xuIiwiLy8gQ29weXJpZ2h0IChDKSAyMDIxIENyZWRpdG9yIENvcnAuIEdyb3VwLlxuLy8gU2VlIExJQ0VOU0UgZm9yIGNvcHlpbmcgaW5mb3JtYXRpb24uXG5cbmltcG9ydCB7IENhcmQgfSBmcm9tICdAL2NhcmQnO1xuaW1wb3J0IHsgTWFya2V0UGxhY2VGb290YmFsbGVyQ2FyZCB9IGZyb20gJ0Bjb21wb25lbnRzL01hcmtldFBsYWNlL01hcmtldFBsYWNlQ2FyZHNHcm91cC9NYXJrZXRQbGFjZUZvb3RiYWxsZXJDYXJkJztcblxuaW1wb3J0ICcuL2luZGV4LnNjc3MnO1xuXG5leHBvcnQgY29uc3QgTWFya2V0UGxhY2VDYXJkc0dyb3VwOiBSZWFjdC5GQzx7IGNhcmRzOiBDYXJkW10gfT4gPSAoeyBjYXJkcyB9KSA9PlxuICAgIDxkaXYgY2xhc3NOYW1lPVwibWFya2V0cGxhY2UtY2FyZHNcIj5cbiAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJtYXJrZXRwbGFjZS1jYXJkc19fd3JhcHBlclwiPlxuICAgICAgICAgICAge2NhcmRzLm1hcCgoY2FyZCwgaW5kZXgpID0+XG4gICAgICAgICAgICAgICAgPE1hcmtldFBsYWNlRm9vdGJhbGxlckNhcmQgY2FyZD17Y2FyZH0ga2V5PXtpbmRleH0gLz5cbiAgICAgICAgICAgICl9XG4gICAgICAgIDwvZGl2PlxuICAgIDwvZGl2PjtcbiIsIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5pbXBvcnQgeyB1c2VEaXNwYXRjaCwgdXNlU2VsZWN0b3IgfSBmcm9tICdyZWFjdC1yZWR1eCc7XG5cbmltcG9ydCB7IE1hcmtldFBsYWNlQ2FyZHNHcm91cCB9IGZyb20gJ0Bjb21wb25lbnRzL01hcmtldFBsYWNlL01hcmtldFBsYWNlQ2FyZHNHcm91cCc7XG5pbXBvcnQgeyBGaWx0ZXJGaWVsZCB9IGZyb20gJ0Bjb21wb25lbnRzL2NvbW1vbi9GaWx0ZXJGaWVsZCc7XG5pbXBvcnQgeyBQYWdpbmF0b3IgfSBmcm9tICdAY29tcG9uZW50cy9jb21tb24vUGFnaW5hdG9yJztcblxuaW1wb3J0IHsgUm9vdFN0YXRlIH0gZnJvbSAnQC9hcHAvc3RvcmUnO1xuaW1wb3J0IHsgY3JlYXRlQ2FyZExpc3QgfSBmcm9tICdAL2FwcC9zdG9yZS9hY3Rpb25zL2NhcmRzJztcblxuaW1wb3J0ICcuL2luZGV4LnNjc3MnO1xuXG5jb25zdCBNYXJrZXRQbGFjZSA9ICgpID0+IHtcbiAgICAvLyBjb25zdCBkaXNwYXRjaCA9IHVzZURpc3BhdGNoKCk7XG4gICAgLy8gZGlzcGF0Y2goY3JlYXRlQ2FyZExpc3QoKSk7XG4gICAgY29uc3QgY2FyZHMgPSB1c2VTZWxlY3Rvcigoc3RhdGU6IFJvb3RTdGF0ZSkgPT4gc3RhdGUuY2FyZHNSZWR1Y2VyLmNhcmRzKTtcblxuICAgIHJldHVybiAoXG4gICAgICAgIDxzZWN0aW9uIGNsYXNzTmFtZT1cIm1hcmtldHBsYWNlXCI+XG4gICAgICAgICAgICA8RmlsdGVyRmllbGRcbiAgICAgICAgICAgICAgICB0aXRsZT1cIk1BUktFVFBMQUNFXCJcbiAgICAgICAgICAgIC8+XG4gICAgICAgICAgICA8TWFya2V0UGxhY2VDYXJkc0dyb3VwXG4gICAgICAgICAgICAgICAgY2FyZHM9e2NhcmRzfVxuICAgICAgICAgICAgLz5cbiAgICAgICAgICAgIDxQYWdpbmF0b3JcbiAgICAgICAgICAgICAgICBpdGVtQ291bnQ9e2NhcmRzLmxlbmd0aH0gLz5cbiAgICAgICAgPC9zZWN0aW9uPlxuICAgICk7XG59O1xuXG5leHBvcnQgZGVmYXVsdCBNYXJrZXRQbGFjZTtcbiJdLCJzb3VyY2VSb290IjoiIn0=