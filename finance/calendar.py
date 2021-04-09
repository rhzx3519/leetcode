import datetime
from yahoo_earnings_calendar import YahooEarningsCalendar

date_from = datetime.datetime.strptime(
    'May 5 2017  10:00AM', '%b %d %Y %I:%M%p')
date_to = datetime.datetime.strptime(
    'May 8 2017  1:00PM', '%b %d %Y %I:%M%p')
yec = YahooEarningsCalendar()
print(yec.earnings_on(date_from))
print(yec.earnings_between(date_from, date_to))