package parsers

import (
	"context"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"

	shared "printraduga_parser/shared"
)

type GcTranslusentParser struct {
}

func (p GcTranslusentParser) Parse() (shared.ParseResult, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true))

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx) // chromedp.WithDebugf(log.Printf),

	defer cancel()
	// ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	// defer cancel()

	link := "https://gcprint.ru/catalog/nakleyki/stikery-na-prozrachnoy-plenke-s-pechatyu-belym-tsvetom/"
	// run task list
	var res string
	err := chromedp.Run(taskCtx,
		chromedp.EmulateViewport(1000, 1200),
		chromedp.Navigate(link),
		chromedp.WaitVisible("div.calc-tabs__item:nth-child(3)"),
		chromedp.Click(".holst-calc__item_type > div:nth-child(1)"),
		chromedp.Click("#typeRound"),
		chromedp.Click("body > div.g-page > section.g-section.g-section_calculation.g-section_calculation-simple > div > div > div.calc-tabs__item.js-tabs-content.is-active > form > div.holst__col.holst__col_left > div > div.holst-calc__row.holst-calc__item.holst-calc__item_material > div"),
		chromedp.Click("body > div.jquery-modal.blocker.current > div > div > div > div:nth-child(2) > div.modal-materials__group-name.js-group-name"),
		chromedp.Click("body > div.jquery-modal.blocker.current > div > div > div > div:nth-child(2) > div.modal-materials__group-wrap.js-group-wrap > div > div.modal-materials__group-name.js-group-name"),
		chromedp.Click("body > div.jquery-modal.blocker.current > div > div > div > div:nth-child(2) > div.modal-materials__group-wrap.js-group-wrap > div > div.modal-materials__group-wrap.js-group-wrap > div:nth-child(2)"),
		chromedp.Evaluate(`document.querySelector("#SIZE_1_DP_PLOTTER").click()`, nil),
		chromedp.Evaluate(`document.querySelector("#colorCmykPlot").click()`, nil),
		chromedp.SetValue("body > div.g-page > section.g-section.g-section_calculation.g-section_calculation-simple > div > div > div.calc-tabs__item.js-tabs-content.is-active > form > div.holst__col.holst__col_left > div > div:nth-child(10) > div > input", "100"),
		chromedp.SendKeys("body > div.g-page > section.g-section.g-section_calculation.g-section_calculation-simple > div > div > div.calc-tabs__item.js-tabs-content.is-active > form > div.holst__col.holst__col_left > div > div:nth-child(10) > div > input", "0"),
		chromedp.Click("body > div.g-page > section.g-section.g-section_calculation.g-section_calculation-simple > div > div > div.calc-tabs__item.js-tabs-content.is-active > form > div.holst__col.holst__col_left > div > p"),
		chromedp.WaitVisible("body > div.g-page > section.g-section.g-section_calculation.g-section_calculation-simple > div > div > div.calc-tabs__item.js-tabs-content.is-active > form > div.holst__col.holst__col_right > div.holst-result > div > div.holst-result__item.holst-result__item_orange > div:nth-child(2) > p > span.js-price"),
		chromedp.Text("body > div.g-page > section.g-section.g-section_calculation.g-section_calculation-simple > div > div > div.calc-tabs__item.js-tabs-content.is-active > form > div.holst__col.holst__col_right > div.holst-result > div > div.holst-result__item.holst-result__item_orange > div:nth-child(2) > p > span.js-price", &res, chromedp.NodeVisible),
		// chromedp.Sleep(time.Hour),
	)
	if err != nil {
		return shared.ParseResult{}, err
	}
	trimmedString := strings.Replace(res, " ", "", -1)
	intVar, err := strconv.Atoi(trimmedString)
	if err != nil {
		return shared.ParseResult{}, err
	}

	return shared.ParseResult{
		ParserType: "Translusent",
		Data: shared.CostData{
			Name: "GCprint",
			Cost: intVar,
			Link: link,
		},
	}, nil
}
