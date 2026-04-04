# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удаленно включить/выключать отопление в своих домах.
- Система поддерживает…
- …

**Мониторинг температуры:**

- Пользователи могут просматривать текущую температуру в своих домах
- Система поддерживает…
- …

### 2. Анализ архитектуры монолитного приложения

Перечислите здесь основные особенности текущего приложения: какой язык программирования используется, какая база данных, как организовано взаимодействие между компонентами и так далее.

1. Проект написан на языке Go версии 1.22. Разрыв с последней в 4 версии или 2 года, стоит держать разрыв в 1-2 версии или 1 год, чтобы не отставать от новых фич языка и при этом не сталкиваться с проблемами несовместимости.
2. База данных PostgreSQL 16 версии. Разрыв с последней в 2 версии или 1 год, не критично т.к. текущую версию ещё будут поддерживать до 2028 года.
3. Архитектура монолитная, все компоненты находятся в одном приложении и взаимодействуют напрямую друг с другом. Масштабирование ограничено, т.к. при увеличении нагрузки на одну часть приложения, приходится масштабировать всё приложение целиком, что может быть неэффективно и дорого. Так же из-за данной архитектуры невозможно выборочно обновлять или изменять отдельные части приложения, что может замедлить разработку и внедрение новых фич.
4. Сервис синхронный, запросы обрабатываются последовательно. Является проблемой т.к. работа приложения сильно зависит от IO задач (работа с базой данных и опрос датчиков), достижение требуемой производительности может потребовать намного больше ресурсов, чем при использовании асинхронного подхода.

### 3. Определение доменов и границы контекстов

Опишите здесь домены, которые вы выделили.

* Домен подключения устройств:
  * поддомен: датчики температуры
    * контекст: подключение датчика температуры
  * поддомен: устройства отопления
  * поддомен: камеры видеонаблюдения
* Домен управление устройствами:
  * поддомен: устройства отопления
    * контекст: изменение температуры
  * поддомен: камеры видеонаблюдения
    * контекст: смена режима работы камеры
* Домен мониторинга:
  * поддомен: датчики температуры
    * контекст: просмотр текущей температуры
  * поддомен: устройства отопления
    * контекст: просмотр текущего состояния устройств отопления
  * поддомен: камеры видеонаблюдения
    * контекст: просмотр текущего состояния камер видеонаблюдения
* Домен управления пользователями:
  * поддомен: пользователи
    * контекст: регистрация и авторизация пользователей
    * контекст: управление правами доступа пользователей
* Домен автоматизации:
  * поддомен: сценарии автоматизации
    * контекст: создание и управление сценариями автоматизации
    * контекст: выполнение сценариев автоматизации

### **4. Проблемы монолитного решения**

- Для релиза новой версии требуется остановка всего приложения, что может привести к простою для пользователей и потере доходов.
- Нельзя масштабировать отдельные части приложения, что может привести к неэффективному использованию ресурсов и увеличению затрат.
- При возникновении ошибки в одной части приложения, может пострадать вся система, что может привести к потере данных и ухудшению пользовательского опыта.
- Сложность внедрения новых фич и обновлений, т.к. изменения в одной части приложения могут повлиять на другие части, что может замедлить разработку и внедрение

Если вы считаете, что текущее решение не вызывает проблем, аргументируйте свою позицию.

### 5. Визуализация контекста системы — диаграмма С4

Добавьте сюда диаграмму контекста в модели C4.

Чтобы добавить ссылку в файл Readme.md, нужно использовать синтаксис Markdown. Это делают так:

```markdown
[Текст ссылки](URL)
```

Замените `Текст ссылки` текстом, который хотите использовать для ссылки. Вместо `URL` вставьте адрес, на который должна вести ссылка. Например:


[Ссылка на диаграмму](https://www.plantuml.com/plantuml/uml/bLBBJjj05DtdAwPP1GdYJLVTLGYfe3uYNDGb6X83aVoIFMxfZZHKNL31lXlKzWSCYi4aZlc5Et_KSoTEd4XXKNandZjtpZbppdtHzxjkqdF4gL9X_Cgo8lcbVT9NfybH4ZjHD_2LuwjsD_2iq9M-IZntArOzjtzaRR_Swd3fuDrRovEnSYAK3hVvnfbgX-XiD1eT9ue3lyLOcw9vWAM9bMwef8d-IGzOxoZLtoZZlQ0dVKuvJL2-HsbAkRuMz0U_JpY_f4T18vmLpOqvSfo4OH0fWDgpFuF_Yvie1S28Kg1YEyAN0VsUnfJ_3v7z5M_pMfzavBq9y12caj9cr0bBrnTK0gnWeZK0BEStu0VcaKKjPuo-A_sZPi-57XfTAJoPBEee5ZG5P64w60Wz9JqQIBXd3kYk8gGCcYlHXpjdgjLpFlUxJZopFTNdk3Hb9_I0dQ1iBUOYfLMTB2Eh5Jr_5qd7Od_7_YXUTtOtnPxzuRteIIN5SH1vKh74VSUjvcBPsdeqm981Ui0M39OC11Eky-r5zYi0Voj5I5zVCyGix99kZQc5lXHc05Pc8ZT6SlCobnvMJ3N1tCsvqE5oWYCpSd16VpSS9M3cWOuv-N4zbsl31jq_)
или
[Путь до файла](./docs/as_is_c4_context.puml)

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

Добавьте диаграмму.
As is
[Текст ссылки](https://www.plantuml.com/plantuml/uml/jLBTRjCm6BtFKtYzePEwLaXSSGSRGg01fM5NQ9gSnDiY9N5akum98QtT41SWU0EGU88mjT3_lOBzHfmlRGD8lLrLoOxzdVETS_njnLcOT2F1Jet9zVrMI_6pvabiJhj1LmiLoio8p3H3cRRoX6UccoOxsaH97BHsx-sqwUZWsNuhvTe8XQEZNKAyDSqSbTxMRR3pE1DgoN_d5XgakSm8KpRmlBPGMcPrOckmsW6M6bRxtIxir7sslx5UsyBrxP2-cRrn5tPafpZ-XQEv7RclpFx0zWRtd-oOsG7kPhLjppz7qxGUYJTnAFIIj70Ne9hR8TQUdO1foDobwxa-Q1TsJ3lgLsaOe6ZglhYFPRCvEX30pghjPYnyQDPvvBzY3zfF6kof_-8vYxZXx5Ygb1keMyaxUYYxMSJLUJsy2zphwW4EcS-vpyWMsNTNTe4sCrjiSdjOssRlF8PVXQ_pF0OXsfEUnx9qzCt-XFKUq4Cu6f5QxFgLKgARKg3N2Bv0mQoCO5YPexHTRwDjr8uO_H4OvOg4-ijJkFrFsbvGZmB8sKCW0TU2fBMbDYTA-2_twSitKY506efbR3oB6C_p90wvYMN6jE56WFUzvtL4ifgUfUE_euDtWZde8ysGtYkdY2oD0IKS8TRpiGPUhjww7Am179JFxZwKtXQWe0aiu7X2Rp6X-CkY0St5-cwwflYtKT2eTGIFD3jMCYqjdQdOmAUHD11_0000)
или
[Путь до файла](./docs/as_is_c4_container.puml)

To be
[Текст ссылки](https://www.plantuml.com/plantuml/uml/pLJRJXD16BxVfnXwOqs0niGhxn0CqO1MAbU6aEdkq6wof-nEnSQOK50nqS0RgF49AhQofNHvXTatylkdtIqzYDwGs6Gxn__xllyq2xPdEWrBiBgKGVIWL0hv_dBDblL6jH69qFGzAJovRFfkAN2u9nkkinH9ox6hfTNxIsNQsjxSABayaYB4rkugMRoMom5k6WktWSvjJuH3_3ktGE06laSsci0moYACVAz8Q8kii8sXW55zLOdgfv_LESOprKuFLPVMJ5sb1ofMVMp_mjP0nUafKzynlSBv9zLZgeEpoy9Tez84gLkqn2lR5D4Gj9qU06gWkXXRwW98NPOUfQtq4B29wj6C-78D7LpgfoVfiIORW84Cr6K-dMyBFoAszQZoWgsK7o-o9vLdJvd59MTxVkXo2NKZURiF_OPdyR1PDE_1xcMUW1rMuRp2o2roFtsd4j1SglOyRtSATzaRW-6NsnTvGzPtUM03TGLJnZsBnB1rBiK-RvBASbFMVK_x-aNzG51QmEt2rvZymEoGWVj81t-0_nWgOWIadNx0_8If6EVg32U9pihv-EHgNlxmzHOPEO7mVyH7q1DccEeC9v02GwOkI3F2DUcRoprU4-78lStXYlnOzs_q9ii8XQrFXNEc6MD8ebHtyoHILQMjkUd5tLXxZiAW-j6HqRipCaDHioCPSkbdKhV6ErC9c2qsgatYIO0_lJyhR8RSz1r7c6FgrMpztyLfhLgqgf2pXWQfyaZeoFYHh8MYydo3LXN8HaFH8p5i8ofc7WIERN9f-nwB9AeQk6yDOqiuezVaMWUCtnCqBpo4I18E_IrTckHN3_Pad9jThp92mAE1gnrKwp6ruLIjZgopvsOwzmNr3SO_Xuom6fRyAKK8aW2gs-E6x1BuTKKPKURRKLOU14fu-hbAt-SrH5hHhTI_b9I4-X2NEiIVdVKpY17XL15DlAKJf5Ih4wHW_BCVUOYHT_G4joAs5_gkJkj-Q4laMlgZkzE1_0UJITL_CJ1-0m00)
или
[Путь до файла](./docs/to_be_c4_container.puml)

**Диаграмма компонентов (Components)**

Добавьте диаграмму для каждого из выделенных микросервисов.
As is
[Текст ссылки](https://www.plantuml.com/plantuml/uml/jLB1RjD04BtxAwO-kL8Q2uaJDsWH0XK8CJsXgbhRszQIlMljRYWLGcfIeHv0y0b8_GAhYY5jclGNPl-8MISnHVI4r1oowvrzxysyjskPIF91ok7Gc58_2aF5Zhe7cJaSn0FDLIA5uS9q4rc4PSw46HJvXZPfdiNXviEDdZlQURaT5amRhkXeRzCXFfXkIJUzjVjWPvdWwjxuYr8AXsduFQ_7bPOYS6mTy7TeyMd57pZ7gVsAlsYzmCgEi7RVW3Vs12SufUEVT3J3YVqEU4xR6wf_mKl0CTMk_NKlEYRH8aZv-opXUa5QsHEYcc5Dwn2lYAa6Usg7TaIqSxnqEwVNzZ2cIrFxpPurOZDIS02yRhTtOsaE3VsdqLjut7lUWHVHwrUGCiDWNwg26UgkDNeh3jAO85jKZfdc3jwJsXme7htPTb-HaOhed4KB4YPIO6MPPmapcHIW3JEydNZlYHo8bAdZ87b0w7UjqgvhPjTVXuyUqA-D8j1SQEdiZveD9AkAKtycRmLokZJ2-KRUL0jFx6V0kHtTuGRWf2ariLNapXLMG7-KK0CUEpDNjUagK1hba-VrURvy5cr_Vvu5mMgI6bBw5EIaKQk9b9G9xQ1Bq2h-VoPmx-Ql2VwpqzjQSSVlSvPgs5UoQ68wKfcXx8oaczvl)
или
[Путь до файла](./docs/as_is_c4_component.puml)

To be
[Текст ссылки](https://www.plantuml.com/plantuml/uml/pLJ1QXin4BtlLoW-ET3O57hgBKrAA3ZWsisfX23lAlR2xchHQd5CASGnfOSKscDFRUaNR5sxSMB7-GNfZ_eahxTEIo440eO5IJFpPj-RaMR3ZzmWzHZfI1adJwjLGRkLjYyxQIjDc71v95aaAouFgt508vc6mLOYlS1lLNUURDMDwM2_LWrf8fc0wMZ784va9KQai_4GsQaVCV5W-SEOHmPZypjEdOc4HsMJxH4f7M0fOLLVrLnDzGTrWNMeCdscnkPCr8q-LHCrXVaNJDTgeZyIzHtR6_ZFrILH8_XcfKsdaBBRAdij9DsbahPemimFwZoHRS6QBslcr9HSi8GO4iH35A6HHsWS1xvBfSyZaaWgsMhESaWZscRYYCO-4kmlJYRN7hS2lgWVQgpxe3NMxy4O8h0x8utdpLTak_x2y9mQxaDep10m0-qHT5pgFjGEhJ3T9tfWpX0zBpmPOYV4V5SenCIK0NPh1KMdsKEnmwDdF8qyAdfbzp7ul2wO7vBjL7QiGaDZ3yAwj7TOsYHl75AqfnpmjXy1M3ChCOqMF0WM4kizsK5HZj5LLCbVQTW0McrNrDW8pB2zrkVw7L5pQFzNYrK1bsrh1crNAZEoq44B7W5J-ARBXaFb2bF1sduYXMqKc3HMZclDEs5kAdpPC_Sltu9lxbq_kh7W9uPhtQphLzYtZjDWGJv7bfTNVL93v9hRcN61o5EWyGhnBuFv0_Zh3gbloFmJyGjs_ySjYw8-swUUpwn5C5xl8WDopVGWn-M8s-FZxl3RH5UkqDrJvCMv5pPdVozi0S8mnYkbNaAE1G_j9Eh2WG4iEM9_0m00)
или
[Путь до файла](./docs/to_be_c4_component_auth.puml)

**Диаграмма кода (Code)**

Добавьте одну диаграмму или несколько.

# Задание 3. Разработка ER-диаграммы

Добавьте сюда ER-диаграмму. Она должна отражать ключевые сущности системы, их атрибуты и тип связей между ними.

[Ссылка](https://www.plantuml.com/plantuml/uml/jLJ1IiD05BplL-nHmHuKww6KqeE88DxgEJopasPnamrlNrA2-EzkareRJUcbpIdCi9kPD_Eo348iQOvIAOTCj7ZDTfsBDfQHYlMHG2bMMlqMmFfJwVXOCzEs53sVVw1TB3eigvrBIYemqq4ukPjEEnW5MJU4SWfvKb486yzcIB5tk_CFznsGCaXW4fAPDl5DR86Eg2ipJXWu_2E-zYwUVHqmj85rD7PeXPobC1H6nZ327z1lId0olENfPNV9sPK_eh4fLcrrIlvl29SQOY_bX0o9IeyD9z-cS8my6F_EG6oUfTYkkdW73pE81POwNDRkAj85XDiLYWeSFyYdp9jkGiSSDkJg657cF4dVUwImpPYqyvdfzdToTNlpUZtfrq3ro_23zQcgWvwDCdjqR2nABL6QxG_b6m00)
[Путь до файла](./docs/er.puml)

# Задание 4. Создание и документирование API

### 1. Тип API

Укажите, какой тип API вы будете использовать для взаимодействия микросервисов. Объясните своё решение.

Будет выбран Rest API т.к.:
1. Широко используется и поддерживается
2. Прост в реализации и понимании
3. Команда уже имеет опыт работы с ним

Предпочтение будет отдано async api т.к.:
1. Приложение является IO-зависимым, и использование асинхронного подхода может улучшить производительность и отзывчивость системы

### 2. Документация API

Здесь приложите ссылки на документацию API для микросервисов, которые вы спроектировали в первой части проектной работы. Для документирования используйте Swagger/OpenAPI или AsyncAPI.

[Путь до файла swagger документация](./docs/swagger.yaml)

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


# **Задание 6. Разработка MVP**

Необходимо создать новые микросервисы и обеспечить их интеграции с существующим монолитом для плавного перехода к микросервисной архитектуре. 

### **Что нужно сделать**

1. Создайте новые микросервисы для управления телеметрией и устройствами (с простейшей логикой), которые будут интегрированы с существующим монолитным приложением. Каждый микросервис на своем ООП языке.
2. Обеспечьте взаимодействие между микросервисами и монолитом (при желании с помощью брокера сообщений), чтобы постепенно перенести функциональность из монолита в микросервисы. 

В результате у вас должны быть созданы Dockerfiles и docker-compose для запуска микросервисов. 