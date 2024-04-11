package icu.aimerny.elix.entity;

import icu.aimerny.elix.Elix;
import icu.aimerny.elix.entity.custom.TigerEntity;
import net.fabricmc.fabric.api.object.builder.v1.entity.FabricDefaultAttributeRegistry;
import net.fabricmc.fabric.api.object.builder.v1.entity.FabricEntityTypeBuilder;
import net.fabricmc.fabric.impl.object.builder.FabricEntityType;
import net.minecraft.entity.EntityDimensions;
import net.minecraft.entity.EntityType;
import net.minecraft.entity.SpawnGroup;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.util.Identifier;

public class ModEntities {

    public static final EntityType<TigerEntity> TIGER = Registry.register(Registries.ENTITY_TYPE,
            new Identifier(Elix.MOD_ID, "tiger"),
            FabricEntityTypeBuilder.create(SpawnGroup.CREATURE, TigerEntity::new)
                    .dimensions(EntityDimensions.fixed(1f,1f)).build()
    );

    public static void init() {
        FabricDefaultAttributeRegistry.register(ModEntities.TIGER, TigerEntity.createMobAttributes());
    }

}
