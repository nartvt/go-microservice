package router

import (
	productRepo "order-service/app/domain/usercases/product/repo"
	bodyRecordRepo "order-service/app/domain/usercases/record/body/repo"
	diaryRecordRepo "order-service/app/domain/usercases/record/diary/repo"
	exerciseRecordRepo "order-service/app/domain/usercases/record/exercise/repo"
	sectionRepo "order-service/app/domain/usercases/section/repo"
	productHandler "order-service/app/transport/product/handler"
	bodyRecordHandler "order-service/app/transport/record/body/handler"
	diaryRecordHandler "order-service/app/transport/record/diary/handler"
	exerciseRecordHandler "order-service/app/transport/record/exercise/handler"
	sectionHandler "order-service/app/transport/section/handler"
)

func setupNewsfeedRoute(route fiber.Router) {

	groupSectionHandler := sectionHandler.SectionHandler{
		SectionDomain: sectionRepo.NewNewsfeedSectionRepo(),
	}
	groupProductHandler := productHandler.ProductHandler{
		ProductDomain: productRepo.NewProductRepo(),
	}

	groupBodyRecordHandler := bodyRecordHandler.BodyRecordHandler{
		BodyRecordDomain: bodyRecordRepo.NewUserBodyRecordRepo(),
	}

	groupDiaryRecordHandler := diaryRecordHandler.DiaryHandler{
		DiaryDomain: diaryRecordRepo.NewUserBodyRecordRepoRepo(),
	}

	groupExerciseRecordHandler := exerciseRecordHandler.ExerciseHandler{
		ExerciseDomain: exerciseRecordRepo.NewUserExerciseRecordRepoRepo(),
	}

	groupNewsfeedSection := route.Group("/newsfeed/sections")
	{
		GET(groupNewsfeedSection, "", groupSectionHandler.GetSections)
	}

	groupProduct := groupNewsfeedSection.Group("product")
	{
		GET(groupProduct, "", groupProductHandler.CreateProduct)
		GET(groupProduct, "/:sectionId", groupProductHandler.GetProductBySectionId)
	}

	groupBodyRecords := route.Group("/body-records")
	{
		GET(groupBodyRecords, "", groupBodyRecordHandler.GetBodyRecordsByUserId)
	}

	groupDiary := route.Group("/diary-records")
	{
		GET(groupDiary, "", groupDiaryRecordHandler.GetDiariesByUserId)
	}

	groupExercise := route.Group("/exercise-records")
	{
		GET(groupExercise, "", groupExerciseRecordHandler.GetExerciseByUserId)
	}
}
